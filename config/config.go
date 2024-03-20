package config

//配置文件的结构体
//@author: [Gaayean]
type Server struct {
	Currencies []Currency `json:"currencies" mapstructure:"currencies" yaml:"currencies"`
	RateApi    string     `json:"rateApi" mapstructure:"rateApi" yaml:"rateApi"`
}

type Currency struct {
	Name string `json:"name" mapstructure:"name" yaml:"name"`
	Sign string `json:"sign" mapstructure:"sign" yaml:"sign"`
}
