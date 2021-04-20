package env

import (
	"fmt"
	"github.com/go-ini/ini"
	"log"
	"time"
)
var (
	Cfg *ini.File

	RunMode string

	HTTPPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration

	PageSize  int
	JwtSecret string
)
func init() {
	var err error
	Cfg, err = ini.Load(".env")
	if err != nil {
		log.Fatalf("Fail to parse '.env': %v", err)
	}
	LoadBase()
}

func LoadBase() {
	RunMode = Cfg.Section("").Key("RUN_MODE").MustString("debug")

	fmt.Println("Current Mode",RunMode)
}

func Load(key,section string) string {
	//if (! function_exists('env')) {
    /**
     * Gets the value of an environment variable.
     *
     * @param  string  $key
     * @param  mixed  $default
     * @return mixed
     */
    //function env($key, $default = null)
    //{
    //    return Env::get($key, $default);
    //}
	// return Env::get($key, $default);

	 return Cfg.Section(section).Key(key).String()
}

func New(){
	fmt.Println("mypackage.New")
}