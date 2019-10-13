package gredis_test

import (
	"errors"
	"fmt"
	"github.com/gogf/gf/database/gredis"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/gfile"
	"github.com/gogf/gf/os/gtime"
	"github.com/gogf/gf/test/gtest"
	"github.com/gogf/gf/util/gconv"
	"os"
	"strings"
	"testing"
)

var (
	Clusterip     = "192.168.0.104" //
	Pass1         = ""       //123456 com:123456 home:"" ci:""
	port          = 6379           //com:8579  home: 6379
	ClustersNodes = []string{Clusterip + ":7001", Clusterip + ":7002", Clusterip + ":7003", Clusterip + ":7004", Clusterip + ":7005", Clusterip + ":7006"}
	config        = gredis.Config{
		Host: Clusterip, //192.168.0.55 127.0.0.1
		Port: port,      //8579 6379
		Db:   1,
		Pass: "", // when is ci,no pass   com yyb513941 home:""
	}
)

func init() {
	gredis.FlagBanCluster = false
	// pwd  = "123456"    when is ci,no pass
	config := `[rediscluster]
    [rediscluster.default]
        host = "` + strings.Join(ClustersNodes, ",") + `"
		pwd  ="` + Pass1 + `"
        
[redis]
     default = "` + Clusterip + `:` + gconv.String(port) + `,1"` // 8579  6379
	err := createTestFile("config.toml", config)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func createTestFile(filename, content string) error {
	//TempDir := testpath()
	err := gfile.PutContents(filename, content)
	return err
}

// get testdir
func testpath() string {
	return os.TempDir()
}





func Test_RedisDo(t *testing.T) {
	gtest.Case(t, func() {
		var(
			err = errors.New("")
			s string
			ss []string
			n int
			n64 int64
			//f64s[]float64
			sss [][]string
			//i interface{}

		)

		redis := gredis.New(config)
		defer redis.Close()
		defer redis.Del("k_1")
		defer redis.Del("k_2")
		defer redis.Del("dlist1")
		defer redis.Del("dlist2")
		defer redis.Del("k11")
		defer redis.Del("k113")
		defer redis.Del("set1")
		defer redis.Del("set2")
		defer redis.Del("set11")
		defer redis.Del("zset1")
		defer redis.Del("zset2")
		defer redis.Del("hlog1")
		defer redis.Del("hlog2")
		defer redis.Del("geo1")



		redis.Set("k1","kv1")
		redis.Set("k2","kv2")
		ss,err=redis.Mget("k1","k2")
		gtest.Assert(err,nil)
		gtest.Assert(ss,[]string{"kv1","kv2"})

		_,err=redis.Mget()
		gtest.AssertNE(err,nil)

		s,err=redis.Mset("k1","kv11","k2","kv22")
		gtest.Assert(err,nil)
		gtest.Assert(s,"OK")
		s,err=redis.Mset()
		gtest.AssertNE(err,nil)


		s,err=redis.Rename("k1","k11")
		gtest.Assert(err,nil)
		s,err=redis.Get("k11")
		gtest.Assert(s,"kv11")

		_,err=redis.Renamenx("k11","k113")
		gtest.Assert(err,nil)
		s,err=redis.Get("k113")
		gtest.Assert(err,nil)
		gtest.Assert(s,"kv11")

		n,err=redis.Msetnx("k_1","kv_1","k_2","kv_2")
		gtest.Assert(err,nil)
		gtest.Assert(n,1)
		n,err=redis.Msetnx("k_1","kv_1","k_3","kv_3")
		gtest.Assert(err,nil)
		gtest.Assert(n,0)



		n64,err=redis.Lpush("dlist1",1)
		gtest.Assert(err,nil)
		gtest.Assert(n64,1)
		redis.Lpush("dlist1",2)
		redis.Lpush("dlist2",3)
		redis.Lpush("dlist2",4)

		s,err=redis.RpoplPush("dlist1","dlist2")
		gtest.Assert(err,nil)
		gtest.Assert(s,"1")



		redis.Lpush("dlist1",6)
		ss,err=redis.BrPoplPush("dlist1","dlist2",1)
		gtest.Assert(err,nil)
		gtest.AssertGT(gconv.Float32(ss[0]),0)


		//=============set
		redis.Sadd("set1","m1","m2")
		redis.Sadd("set11","m11","m21","m1")
		n,err=redis.Smove("set1","set2","m2")
		gtest.Assert(err,nil)
		gtest.Assert(n,1)


		ss,err=redis.Sinter("set1","set11")
		gtest.Assert(err,nil)
		gtest.Assert(ss[0],"m1")
		ss,err=redis.Sinter()
		gtest.AssertNE(err,nil)


		n64,err=redis.SinterStore("set11","set1")
		gtest.Assert(err,nil)
		gtest.AssertGT(n64,0)

		ss,err=redis.Sunion("set1","set11")
		gtest.Assert(err,nil)
		gtest.Assert(ss[0],"m1")

		n64,err=redis.SunionStore("set11","set1")
		gtest.Assert(err,nil)
		gtest.Assert(n64,1)
		//SunionStore
		redis.Sadd("set1","m1","m2")
		ss,err=redis.Sdiff("set1","set11")
		gtest.Assert(err,nil)
		gtest.Assert(ss[0],"m2")
		n64,err=redis.SdiffStore("set11","set1")
		gtest.Assert(err,nil)
		gtest.Assert(n64,2)

		redis.Zadd("zset1",1,"m1")
		redis.Zadd("zset1",2,"m2")
		n64,err=redis.ZunionStore("zset2",1,"zset1")
		gtest.Assert(err,nil)
		gtest.Assert(n64,2)
		n64,err=redis.ZunionStore("zset2")
		gtest.AssertNE(err,nil)


		n64,err=redis.ZinterStore("zset2",1,"zset1")
		gtest.Assert(err,nil)
		gtest.Assert(n64,2)

		//=========================HyperLogLog
		n,err=redis.PfAdd("hlog1","e1","e2")
		gtest.Assert(err,nil)
		gtest.Assert(n,1)

		n64,err=redis.PfCount("hlog1")
		gtest.Assert(err,nil)
		gtest.Assert(n64,2)

		redis.PfAdd("hlog2","f1","f2")
		s,err=redis.PfMerge("hlog3","hlog1","hlog2")
		gtest.Assert(err,nil)
		gtest.Assert(s,"OK")
		s,err=redis.PfMerge("hlog3")
		gtest.AssertNE(err,nil)

		//=======================================geo
		n,err=redis.GeoAdd("geo1","13.361389", "38.115556", "beijin", "15.087269" ,"37.502669" ,"chengdu")
		gtest.Assert(err,nil)
		gtest.Assert(n,2)

		sss,err2:=redis.GeoPos("geo1","beijin","chengdu")
		gtest.Assert(err2,nil)
		gtest.Assert(len(sss),2)
		gtest.Assert(len(sss[0]),2)

		s,err=redis.GeoDist("geo1","beijin","chengdu","km")
		gtest.Assert(err,nil)
		gtest.AssertGT(gconv.Float64(s),166.1)

		sss,err=redis.GeoRadius("geo1","15","37","200","km","WITHDIST")
		gtest.Assert(err2,nil)
		t.Error(sss)

	})
}


func Test_Clustersg(t *testing.T) {
	gtest.Case(t, func() {
		var(
			n int
			n64 int64
			n64_2 int64
			rr interface{}
			rrs []interface{}
			err = errors.New("")
			s string
			//f64 float64
			ss []string
		)

		gredis.FlagBanCluster = false

		rdb:=g.Redis()

		defer rdb.Del("hash1")
		defer rdb.Del("tn1")
		defer rdb.Del("tn2")
		defer rdb.Del("jjname1_11")
		defer rdb.Del("list1")
		defer rdb.Del("set1")
		defer rdb.Del("zset1")
		defer rdb.Del("hlog1")
		defer rdb.Del("geo1")

		rr, err = rdb.Cluster("info")
		gtest.Assert(err, nil)
		str1 := gconv.String(rr)
		if !strings.Contains(str1, "cluster_state:ok") {
			t.Errorf("cluster errs.")
		}

		//rdb.Del("hash1")

		_, err = rdb.Set("jjname1", "jjqrr1")
		gtest.Assert(err, nil)
		_, err = rdb.Set("jjname2", "jjqrr2")
		_, err = rdb.Set("jjname3", "jjqrr3")
		gtest.Assert(err, nil)
		rr, err2 := rdb.Get("jjname2")
		gtest.Assert(err2, nil)
		gtest.Assert(gconv.String(rr), "jjqrr2")

		rdb.Set("n1","10")

		rr, err= rdb.Get("jjname1")
		gtest.Assert(err, nil)
		gtest.Assert(gconv.String(rr), "jjqrr1")

		rr3, err3 :=rdb.Get("jjname3")
		gtest.Assert(err3, nil)
		gtest.Assert(gconv.String(rr3), "jjqrr3")

		n,_=rdb.Exists("jjname3")
		gtest.Assert(n,1)

		n64,_=rdb.Expire("jjname3",300)
		gtest.Assert(n,1)
		n64 ,_=rdb.Ttl("jjname3")
		gtest.AssertGT(n64,200)

		rr,_=rdb.Dump("jjname3")
		gtest.AssertNE(rr,nil)

		n,_=g.Redis().Expireat("jjname3",gtime.Now().Second()+120)
		gtest.Assert(n,1)

		rrs,_=rdb.Keys("*jjname*")
		gtest.AssertGT(len(rrs),0)

		rr,_=rdb.Object("REFCOUNT","jjname3")
		gtest.AssertGT(gconv.Int(rr),0)

		n,_=rdb.Persist("jjname3")
		gtest.Assert(n,1)
		n,_=rdb.Persist("jjname3_")
		gtest.Assert(n,0)

		n64,_=rdb.Pttl("jjname3")
		gtest.Assert(n64,-1)
		n64,_=rdb.Pttl("jjname3_")
		gtest.AssertLT(n64,0)
		g.Redis().Expire("jjname3",10)
		n64,_=rdb.Pttl("jjname3")
		gtest.AssertGT(n64,5)


		rr,_=rdb.RandomKey()
		gtest.AssertNE(rr,nil)






		_,err=rdb.ReStore("jjname2",100000,"servals")
		gtest.AssertNE(err,nil)

		s,err=rdb.Dump("jjname1")
		gtest.Assert(err,nil)

		s,err=rdb.ReStore("jjname1",100000,s,"replace")
		gtest.Assert(err,nil)
		gtest.Assert(s,"OK")


		n64,err=rdb.Lpush("numlist2",1,3)
		gtest.Assert(err,nil)
		gtest.AssertGT(n64,0)
		n64,err=rdb.Lpush("numlist2",2)

		rrs,err= rdb.Sort("numlist2","desc")
		gtest.Assert(err,nil)
		gtest.Assert(gconv.SliceStr(rrs),[]string{"3","2","1"})


		//=============================del this lists after test
		n,err=rdb.Del("numlist2")
		gtest.Assert(n,1)

		s,err=rdb.Get("numlist2")
		gtest.Assert(err,nil)
		gtest.Assert(s,"")
		// Sort

		rdb.Set("jname2","a")
		n64,err=rdb.Append("jname2","q")
		s,err=rdb.Get("jname2")
		gtest.Assert(s,"aq")

		s,err= rdb.Type("jname2")
		gtest.Assert(err,nil)
		gtest.Assert(s,"string")

		n,err=rdb.Setbit("jname2",3,1)
		gtest.Assert(err,nil)
		gtest.Assert(n,0)

		n,err=rdb.Getbit("jname2",3)
		gtest.Assert(err,nil)
		gtest.Assert(n,1)

		n,err=rdb.BitCount("jname2")
		gtest.Assert(err,nil)
		gtest.Assert(n,8)

		rdb.Set("jname22","tt22")
		n,err=rdb.Setbit("jname22",3,1)
		n,err=rdb.BiTop("and","and-result","jname2","jname22")
		gtest.AssertNE(err,nil)

		n,err=rdb.BitPos("jname22",1)
		gtest.Assert(err,nil)
		gtest.Assert(n,1)


		rrs,err=rdb.BitField("jname2 set a")
		gtest.Assert(err,nil)


		n64,err=rdb.Decr("n1")
		gtest.Assert(err,nil)
		gtest.Assert(n64,9)


		n64,err=rdb.Decrby("n1",2)
		gtest.Assert(err,nil)
		gtest.Assert(n64,7)

		n64,err=rdb.Decrby("n_21",2)
		gtest.Assert(err,nil)
		gtest.Assert(n64,-2)
		n,err=rdb.Del("n_21")
		gtest.Assert(err,nil)

		s,err=rdb.GetRange("jjname1",1,2)
		gtest.Assert(err,nil)
		gtest.Assert(s,"jq")

		rr,err=rdb.GetSet("jjname1","imjjname1")
		gtest.Assert(err,nil)
		gtest.Assert(rr,"jjqrr1")
		s,err=rdb.Get("jjname1")
		gtest.Assert(s,"imjjname1")


		n64,err=rdb.Incr("tn1")
		gtest.Assert(err,nil)
		gtest.Assert(n64,1)
		n64,err=rdb.Incr("tn1")
		gtest.Assert(n64,2)



		n64,err=rdb.IncrBy("tn1",2)
		gtest.Assert(err,nil)
		gtest.Assert(n64,4)


		s,err=rdb.IncrByFloat("tn2",3.4)
		gtest.Assert(err,nil)
		gtest.Assert(s,"3.4")


		s,err=rdb.Psetex("jjname1",1000,"newj1")
		gtest.Assert(err,nil)
		gtest.Assert(s,"OK")

		s,err=rdb.Setex("jjname1",5,"newj11")
		gtest.Assert(err,nil)
		gtest.Assert(s,"OK")

		n,err=rdb.Setnx("jjname1","nn1")
		gtest.Assert(err,nil)
		gtest.Assert(n,0)
		n,err=rdb.Setnx("jjname1_11","nn1_11")
		gtest.Assert(err,nil)
		gtest.Assert(n,1)


		n,err=rdb.SetRange("jjname1",1,"y")
		gtest.Assert(err,nil)
		gtest.Assert(n,6)

		n,err=rdb.Strlen("jjname1")
		gtest.Assert(n,6)


		n,err=rdb.Hset("hash1","field1","v1")
		gtest.Assert(err,nil)
		gtest.Assert(n,1)
		n,err=rdb.Hset("hash1","field1","v1_1")
		gtest.Assert(err,nil)
		gtest.Assert(n,0)

		n,err=rdb.Hsetnx("hash1","field2","v2")
		gtest.Assert(err,nil)
		gtest.Assert(n,1)
		n,err=rdb.Hsetnx("hash1","field1","v11")
		gtest.Assert(err,nil)
		gtest.Assert(n,0)

		s,err=rdb.Hget("hash1","field1")
		gtest.Assert(err,nil)
		gtest.Assert(s,"v1_1")

		n,err=rdb.Hexists("hash1","field1")
		gtest.Assert(err,nil)
		gtest.Assert(n,1)
		n,err=rdb.Hexists("hash1","field111")
		gtest.Assert(err,nil)
		gtest.Assert(n,0)

		n,err=rdb.Hdel("hash1","field1")
		gtest.Assert(err,nil)
		gtest.Assert(n,1)
		n,err=rdb.Hdel("hash1ss","field1")
		gtest.Assert(n,0)

		n,_=rdb.Hlen("hash1")
		gtest.Assert(n,1)

		n,_=rdb.Hstrlen("hash1","field2")
		gtest.Assert(n,2)

		n64,err=rdb.HincrBy("hash1","nums",2)
		gtest.Assert(err,nil)
		gtest.Assert(n64,2)


		s,_=rdb.HincrByFloat("hash1","f1",0.34)
		gtest.Assert(s,"0.34")

		s,err=rdb.Hmset("hash1","mk1","mv1","mk2","mv2")
		gtest.Assert(err,nil)
		gtest.Assert(s,"OK")
		s,_=rdb.Hget("hash1","mk2")
		gtest.Assert(s,"mv2")

		ss,err=rdb.Hmget("hash1", "mk1","mk2")
		gtest.Assert(err,nil)
		gtest.Assert(ss,[]string{"mv1","mv2"})

		ss,err=rdb.Hkeys("hash1")
		gtest.Assert(err,nil)
		gtest.Assert(ss,[]string{"field2","nums","f1","mk1","mk2"})

		ss,err=rdb.Hvals("hash1")
		gtest.Assert(err,nil)
		gtest.Assert(ss,[]string{"v2","2","0.34","mv1","mv2"})

		ss,err=rdb.HgetAll("hash1")
		gtest.Assert(err,nil)
		gtest.Assert(ss,[]string{"field2","v2","nums","2","f1","0.34","mk1","mv1","mk2","mv2"})


		// lists

		n64,_=rdb.Lpush("list1",1)
		n64,err=rdb.Lpushx("list1",3)
		gtest.Assert(err,nil)
		gtest.Assert(n64,2)

		n64,err=rdb.Rpush("list1",4)
		gtest.Assert(err,nil)
		gtest.Assert(n64,3)

		n64,err=rdb.Rpushx("list1",5)
		gtest.Assert(err,nil)
		gtest.Assert(n64,4)

		s,err=rdb.Lpop("list1")
		gtest.Assert(err,nil)
		gtest.Assert(s,"3")

		s,err=rdb.Rpop("list1")
		gtest.Assert(err,nil)
		gtest.Assert(s,"5")


		n64,err=rdb.Lrem("list1",2,4)
		gtest.Assert(err,nil)
		gtest.Assert(n64,1)

		n64,err=rdb.Llen("list1")
		gtest.Assert(err,nil)
		gtest.Assert(n64,1)

		rdb.Lpush("list1","a")
		s,err=rdb.Lindex("list1",1)
		gtest.Assert(err,nil)
		gtest.Assert(s,"1")

		n64,err=rdb.Linsert("list1","BEFORE","a","b")
		gtest.Assert(err,nil)
		gtest.Assert(n64,3)

		s,err=rdb.Lset("list1",1,"c")
		gtest.Assert(err,nil)
		gtest.Assert(s,"OK")

		ss,err=rdb.Lrange("list1",0,2)
		gtest.Assert(err,nil)
		gtest.Assert(ss,[]string{"b","c","1"})

		ss,err=rdb.BlPop("list1",1)
		gtest.Assert(err,nil)
		gtest.Assert(ss[0],"list1")
		gtest.Assert(ss[1],"b")

		ss,err=rdb.BrPop("list1",1)
		gtest.Assert(err,nil)
		gtest.Assert(ss[0],"list1")
		gtest.Assert(ss[1],"1")

		//=============================set
		n64,err=rdb.Sadd("set1","m1")
		gtest.Assert(err,nil)
		gtest.Assert(n64,1)


		n,err=rdb.SisMember("set1","m1")
		gtest.Assert(err,nil)
		gtest.Assert(n,1)

		n,err=rdb.SisMember("set1","m2")
		gtest.Assert(err,nil)
		gtest.Assert(n,0)

		rdb.Sadd("set1","m2")
		rdb.Sadd("set1","m3")
		rdb.Sadd("set1","m4")
		s,err=rdb.Spop("set1")
		gtest.Assert(err,nil)
		gtest.AssertNE(s,"")

		ss,err=rdb.SrandMember("set1",2)
		gtest.Assert(err,nil)
		gtest.Assert(len(ss),2)
		ss,err=rdb.SrandMember("set1")
		gtest.Assert(err,nil)
		gtest.Assert(len(ss),1)

		n64,err=rdb.Scard("set1")
		n,err=rdb.Srem("set1","m2")
		n64_2,_=rdb.Scard("set1")
		gtest.Assert(err,nil)
		gtest.AssertGE(n64,n64_2)


		ss,err=rdb.Smembers("set1")
		gtest.Assert(err,nil)
		gtest.AssertGE(len(ss),0)

		//======================zset
		n,err=rdb.Zadd("zset1",1,"m1")
		gtest.Assert(err,nil)
		gtest.Assert(n,1)
		n,err=rdb.Zadd("zset1",1.1,"m2")
		gtest.Assert(err,nil)
		gtest.Assert(n,1)

		s,err=rdb.Zscore("zset1","m1")
		gtest.Assert(err,nil)
		gtest.Assert(s,"1")


		s,err=rdb.ZinCrby("zset1",1.1,"m1")
		gtest.Assert(err,nil)
		gtest.AssertGT(gconv.Float64(s),2.0)


		n64,err=rdb.Zcard("zset1")
		gtest.Assert(err,nil)
		gtest.Assert(n64,2)

		n64,err=rdb.Zcount("zset1",1,3)
		gtest.Assert(err,nil)
		gtest.Assert(n64,2)

		// Zrange
		ss,err=rdb.Zrange("zset1",0,3)
		gtest.Assert(err,nil)
		gtest.Assert(len(ss),2)
		ss,err=rdb.Zrange("zset1",0,3,"WITHSCORES")
		gtest.Assert(err,nil)
			gtest.Assert(len(ss),4)

		//ZrevRange
		ss,err=rdb.ZrevRange("zset1",0,3)
		gtest.Assert(err,nil)
		gtest.Assert(len(ss),2)
		ss,err=rdb.ZrevRange("zset1",0,3,"WITHSCORES")
		gtest.Assert(err,nil)
		gtest.Assert(len(ss),4)

		// ZRANGEBYSCORE
		ss,err=rdb.ZrangeByScore("zset1","1.1","3.1")
		gtest.Assert(err,nil)
		gtest.Assert(ss,[]string{"m2","m1"})

		ss,err=rdb.ZrevRangeByScore("zset1","3.1","1.1")
		gtest.Assert(err,nil)
		gtest.Assert(ss,[]string{"m1","m2"})

		n64,err=rdb.Zrank("zset1","m1")
		gtest.Assert(err,nil)
		gtest.Assert(n64,1)

		n64,err=rdb.ZrevRank("zset1","m1")
		gtest.Assert(err,nil)
		gtest.Assert(n64,0)

		_,err=rdb.Zrem("zset1")
		gtest.AssertNE(err,nil)

		n,err=rdb.Zrem("zset1","m1")
		gtest.Assert(err,nil)
		gtest.Assert(n,1)

		n64,err=rdb.ZremRangeByRank("zset1",0,5)
		gtest.Assert(err,nil)
		gtest.Assert(n64,1)

		rdb.Zadd("zset1",1,"m1")
		rdb.Zadd("zset1",2,"m2")

		n64,err=rdb.ZremRangeByScore("zset1",0.1,3.1)
		gtest.Assert(err,nil)
		gtest.Assert(n64,2)

		rdb.Zadd("zset1",1,"a1")
		rdb.Zadd("zset1",1,"b1")
		rdb.Zadd("zset1",1,"c1")

		ss,err=rdb.ZrangeByLex("zset1","-","[c")
		gtest.Assert(err,nil)
		gtest.Assert(ss,[]string{"a1","b1"})

		n64,err=rdb.ZlexCount("zset1","-","+")
		gtest.Assert(err,nil)
		gtest.Assert(n64,3)

		n64,err=rdb.ZremRangeByLex("zset1","-","[c")
		gtest.Assert(err,nil)
		gtest.Assert(n64,2)




	})
}

