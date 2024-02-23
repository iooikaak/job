package config

import (
	"flag"
	"github.com/iooikaak/frame/cache/redis/v8"
	"github.com/iooikaak/frame/common/kdniao"
	"github.com/iooikaak/frame/json"
	"github.com/iooikaak/job/common/model/enum"
	"os"
	"path/filepath"

	"github.com/iooikaak/frame/conf/paladin"
	"github.com/iooikaak/frame/conf/paladin/apollo"
	"github.com/iooikaak/frame/elasticsearch"
	"github.com/iooikaak/frame/gins"
	"github.com/iooikaak/frame/gorm"
	"github.com/iooikaak/frame/micro"
	rocketmqConfig "github.com/iooikaak/frame/mq/rocketmq/config"
	bm "github.com/iooikaak/frame/net/http/blademaster"
	"github.com/iooikaak/frame/tracer"
	"github.com/iooikaak/frame/xlog"
)

var (
	Conf     = &Config{}
	confPath string
	M        paladin.YAML
)

// Config struct
type Config struct {
	Log            *xlog.Config                   `yaml:"log" json:"log"`
	Tracer         *tracer.TracingConfig          `yaml:"tracer" json:"tracer"`
	GinServer      *gins.Config                   `yaml:"ginServer" json:"ginServer"`
	MicroServer    *micro.Options                 `yaml:"microServer" json:"microServer"`
	DB             *gorm.Config                   `yaml:"db" json:"db"`
	Redis          *redis.Config                  `yaml:"redis" json:"redis"`
	HttpClient     *bm.ClientConfig               `yaml:"httpClient" json:"httpClient"`
	ServiceName    string                         `yaml:"serviceName" json:"serviceName"`
	Elastic        *elasticsearch.ElasticConfig   `yaml:"elastic" json:"elastic"`
	RocketMq       *rocketmqConfig.RocketmqConfig `yaml:"rocketMq" json:"rocketMq"`
	ExpressTracing *kdniao.Config                 `yaml:"expressTracing" json:"expressTracing"`
	PxqUrl         string                         `yaml:"pxqUrl" json:"pxqUrl"`
}

func init() {
	flag.StringVar(&confPath, "Config", "", "conf values")
	//Init()
}

func Init() (err error) {
	var (
		jsonFIle string
	)
	if confPath != "" {
		jsonFIle, err = filepath.Abs(confPath)
	} else {
		jsonFIle, err = filepath.Abs(enum.JobJson.String())
	}
	if err != nil {
		return
	}
	jsonRead, err := os.ReadFile(jsonFIle)
	if err != nil {
		return
	}
	err = json.Unmarshal(jsonRead, Conf)
	if err != nil {
		return
	}
	return
}

// ApolloInit apollo配置初始化，依赖本地环境变量
func ApolloInit() error {
	//apolloAppName := "dramaJob"
	//_ = os.Setenv("APOLLO_NAMESPACES", apolloAppName)
	//if err := paladin.Init(apollo.PaladinDriverApollo); err != nil {
	//panic(err)
	//}

	//if err := paladin.Get(apolloAppName).UnmarshalYAML(&Conf); err != nil {
	//panic(err)
	//}

	//if err := paladin.Watch(apolloAppName, &M); err != nil {
	//panic(err)
	//}

	//return nil
	_ = os.Setenv(enum.ApolloNamespaces.String(), enum.ApolloAppName.String())
	_ = os.Setenv(enum.ApolloAppID.String(), enum.ServiceName.String())
	_ = os.Setenv(enum.ApolloCluster.String(), enum.ApolloClusterValue.String())
	_ = os.Setenv(enum.ApolloMetaAddr.String(), enum.ApolloMetaAddrValue.String())
	_ = os.Setenv(enum.ApolloCacheDir.String(), enum.ApolloCacheDirValue.String())
	if err := paladin.Init(apollo.PaladinDriverApollo); err != nil {
		panic(err)
	}
	if err := paladin.Get(enum.ApolloAppName.String()).UnmarshalJSON(&Conf); err != nil {
		panic(err)
	}
	if err := paladin.Watch(enum.ApolloAppName.String(), &M); err != nil {
		panic(err)
	}
	return nil
}

func GetPxqUrl() (string, error) {
	return M.Get("pxq_url").String()
}
