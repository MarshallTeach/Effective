package setting

import (
	"github.com/go-ini/ini"
	"log"
	"time"
)

type App struct {
	PrefixUrl string
	JwtSecret string
	PageSize int
	RuntimeRootPath string
	LogSavePath	string
	LogSaveName string
	LogFileExt	string
	TimeFormat	string

	ImageSavePath string
	ImageMaxSize int
	ImageAllowExts []string
}

var AppSetting = &App{}

type Server struct {
	RunMode string
	HttpPort int
	ReadTimeOut time.Duration
	ReadHeaderTimeOut time.Duration
	WriteTimeOut time.Duration
}

var ServerSetting = &Server{}

type Database struct {
	Type string
	User string
	Password string
	Host string
	Name string
	TablePrefix string
}

var DatabaseSetting = &Database{}

type Redis struct {
	Host string
	Password string
	MaxIdle int
	MaxActive int
	IdleTimeout time.Duration
}

var RedisSetting = &Redis{}

var cfg *ini.File

func Setup()  {
	var err error
	cfg, err = ini.Load("conf/app.ini")
	if err != nil {
		log.Fatalf("Fail to parse 'conf/app.ini': %v", err)
	}

	mapTo("app", AppSetting)
	mapTo("server", ServerSetting)
	mapTo("database", DatabaseSetting)

	ServerSetting.ReadTimeOut = ServerSetting.ReadTimeOut * time.Second
	ServerSetting.WriteTimeOut = ServerSetting.WriteTimeOut * time.Second
	ServerSetting.ReadHeaderTimeOut = ServerSetting.ReadHeaderTimeOut * time.Second
}

// mapTo map section
func mapTo(section string, v interface{}) {
	err := cfg.Section(section).MapTo(v)
	if err != nil {
		log.Fatalf("Cfg.MapTo %s err: %v", section, err)
	}
}
