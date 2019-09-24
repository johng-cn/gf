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
	Clusterip     = "192.168.0.55" //
	Pass1         = "123456"       //123456
	port          = 8579           //8579 6379
	ClustersNodes = []string{Clusterip + ":7001", Clusterip + ":7002", Clusterip + ":7003", Clusterip + ":7004", Clusterip + ":7005", Clusterip + ":7006"}
	config        = gredis.Config{
		Host: Clusterip, //192.168.0.55 127.0.0.1
		Port: port,      //8579 6379
		Db:   1,
		Pass: "yyb513941", // when is ci,no pass
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

func Test_ClusterDo(t *testing.T) {
	gtest.Case(t, func() {
		redis := gredis.NewClusterClient(&gredis.ClusterOption{
			Nodes: ClustersNodes,
			Pwd:   Pass1,
		})
		redis.Set("jname2", "jqrr2")
		r, err := redis.Get("jname2")
		gtest.Assert(err, nil)
		gtest.Assert(gconv.String(r), "jqrr2")
	})
}

func Test_Clustersg(t *testing.T) {
	gtest.Case(t, func() {
		var(
			n int
			n64 int64
			rr interface{}
			rrs []interface{}
			err = errors.New("")
			s string
		)

		gredis.FlagBanCluster = false

		_, err = g.Redis().Set("jjname1", "jjqrr1")
		gtest.Assert(err, nil)
		_, err = g.Redis().Set("jjname2", "jjqrr2")
		_, err = g.Redis().Set("jjname3", "jjqrr3")
		gtest.Assert(err, nil)
		rr, err2 := g.Redis().Get("jjname2")
		gtest.Assert(err2, nil)
		gtest.Assert(gconv.String(rr), "jjqrr2")

		rr, err= g.Redis().Get("jjname1")
		gtest.Assert(err, nil)
		gtest.Assert(gconv.String(rr), "jjqrr1")

		rr3, err3 := g.Redis().Get("jjname3")
		gtest.Assert(err3, nil)
		gtest.Assert(gconv.String(rr3), "jjqrr3")

		n,_=g.Redis().Exists("jjname3")
		gtest.Assert(n,1)

		n64,_=g.Redis().Expire("jjname3",300)
		gtest.Assert(n,1)
		n64 ,_=g.Redis().Ttl("jjname3")
		gtest.AssertGT(n64,200)

		rr,_=g.Redis().Dump("jjname3")
		gtest.AssertNE(rr,nil)

		n,_=g.Redis().Expireat("jjname3",gtime.Now().Second()+120)
		gtest.Assert(n,1)

		rrs,_=g.Redis().Keys("*jjname*")
		gtest.AssertGT(len(rrs),0)

		rr,_=g.Redis().Object("REFCOUNT","jjname3")
		gtest.AssertGT(gconv.Int(rr),0)

		n,_=g.Redis().Persist("jjname3")
		gtest.Assert(n,1)
		n,_=g.Redis().Persist("jjname3_")
		gtest.Assert(n,0)

		n64,_=g.Redis().Pttl("jjname3")
		gtest.Assert(n64,-1)
		n64,_=g.Redis().Pttl("jjname3_")
		gtest.AssertLT(n64,0)
		g.Redis().Expire("jjname3",10)
		n64,_=g.Redis().Pttl("jjname3")
		gtest.AssertGT(n64,5)


		rr,_=g.Redis().RandomKey()
		gtest.AssertNE(rr,nil)

		// @todo: rename在集群模式下会报错。
		s,err=g.Redis().Rename("jjname2","{jjname22}")
		gtest.Assert(err,nil)
		gtest.Assert(s,"ok")
		_,err=g.Redis().Rename("jjname2_","jjname22")
		gtest.AssertNE(err,nil)

		n,err=g.Redis().Renamenx("jjname22","jjname1")
		gtest.Assert(n,0)
		gtest.Assert(err,nil)

		n,_=g.Redis().Renamenx("jjname22","jjname222")
		gtest.Assert(n,1)




		// Renamenx


		rr, err = g.Redis().Cluster("info")
		gtest.Assert(err, nil)
		str1 := gconv.String(rr)
		if !strings.Contains(str1, "cluster_state:ok") {
			t.Errorf("cluster errs.")
		}

	})
}