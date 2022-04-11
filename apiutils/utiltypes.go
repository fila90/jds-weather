package apiutils

type TypeLocation struct {
	Country   string  `json:"country"`
	Lat       float32 `json:"lat"`
	Localtime string  `json:"localtime"`
	Lon       float32 `json:"lon"`
	Name      string  `json:"name"`
	Region    string  `json:"region"`
	TimeZone  string  `json:"tz_id"`
}

type TypeCurrentConditioin struct {
	Code int    `json:"code"`
	Icon string `json:"icon"`
	Text string `json:"text"`
}

type TypeCurrent struct {
	Cloud       int                   `json:"cloud"`
	Condition   TypeCurrentConditioin `json:"condition"`
	Feelslike_C float32               `json:"feelslike_c"`
	Feelslike_F float32               `json:"feelslike_f"`
	Gust_Kph    float32               `json:"gust_kph"`
	Gust_Mph    float32               `json:"gust_mph"`
	Humidity    int                   `json:"humidity"`
	Isday       int                   `json:"is_day"`
	Lastupdated string                `json:"last_updated"`
	Precip_In   float32               `json:"precip_in"`
	Precip_Mm   float32               `json:"precip_mm"`
	Pressure_In float32               `json:"pressure_in"`
	Pressure_Mb float32               `json:"pressure_mb"`
	Temp_C      float32               `json:"temp_c"`
	Temp_F      float32               `json:"temp_f"`
	Uv          float32               `json:"uv"`
	Vis_Km      float32               `json:"vis_km"`
	Vis_Miles   float32               `json:"vis_miles"`
	Wind_Degree int                   `json:"wind_degree"`
	Wind_Dir    string                `json:"wind_dir"`
	Wind_Kph    float32               `json:"wind_kph"`
	Wind_Mph    float32               `json:"wind_mph"`
}

type TypeDay struct {
	Avghumidity          float32               `json:"avghumidity"`
	Avgtemp_C            float32               `json:"avgtemp_c"`
	Avgtemp_F            float32               `json:"avgtemp_f"`
	Avgvis_Km            float32               `json:"avgvis_km"`
	Avgvis_Miles         float32               `json:"avgvis_miles"`
	Condition            TypeCurrentConditioin `json:"condition"`
	Daily_Chance_Of_Rain int                   `json:"daily_chance_of_rain"`
	Daily_Chance_Of_Snow int                   `json:"daily_chance_of_snow"`
	Daily_Will_It_Rain   int                   `json:"daily_will_it_rain"`
	Daily_Will_It_Snow   int                   `json:"daily_will_it_snow"`
	Maxtemp_C            float32               `json:"maxtemp_c"`
	Maxtemp_F            float32               `json:"maxtemp_f"`
	Maxwind_Kph          float32               `json:"maxwind_kph"`
	Maxwind_Mph          float32               `json:"maxwind_mph"`
	Mintemp_C            float32               `json:"mintemp_c"`
	Mintemp_F            float32               `json:"mintemp_f"`
	Totalprecip_In       float32               `json:"totalprecip_in"`
	Totalprecip_Mm       float32               `json:"totalprecip_mm"`
	Uv                   float32               `json:"uv"`
}

type TypeHour struct {
	Chance_Of_Rain int                   `json:"chance_of_rain"`
	Chance_Of_Snow int                   `json:"chance_of_snow"`
	Cloud          int                   `json:"cloud"`
	Condition      TypeCurrentConditioin `json:"conditoin"`
	Dewpoint_C     float32               `json:"dewpoint_c"`
	Dewpoint_F     float32               `json:"dewpoint_f"`
	Feelslike_C    float32               `json:"feelslike_c"`
	Feelslike_F    float32               `json:"feelslike_f"`
	Gust_Kph       float32               `json:"gust_kph"`
	Gust_Mph       float32               `json:"gust_mph"`
	Heatindex_C    float32               `json:"heatindex_c"`
	Heatindex_F    float32               `json:"heatindex_f"`
	Humidity       int                   `json:"humidity"`
	Is_Day         int                   `json:"is_day"`
	Precip_In      float32               `json:"precip_in"`
	Precip_Mm      float32               `json:"precip_mm"`
	Pressure_In    float32               `json:"pressure_in"`
	Pressure_Mb    float32               `json:"pressure_mb"`
	Temp_C         float32               `json:"temp_c"`
	Temp_F         float32               `json:"temp_f"`
	Time           string                `json:"time"`
	Uv             float32               `json:"uv"`
	Vis_Km         float32               `json:"vis_km"`
	Vis_Miles      float32               `json:"vis_miles"`
	Will_It_Rain   int                   `json:"will_it_rain"`
	Will_It_Snow   int                   `json:"will_it_snow"`
	Wind_Degree    int                   `json:"wind_degree"`
	Wind_Dir       string                `json:"wind_dir"`
	Wind_Kph       float32               `json:"wind_kph"`
	Wind_Mph       float32               `json:"wind_mph"`
	Windchill_C    float32               `json:"windchill_c"`
	Windchill_F    float32               `json:"windchill_f"`
}

type TypeSearch struct {
	Country string  `json:"country"`
	Id      int     `json:"id"`
	Lat     float32 `json:"lat"`
	Lon     float32 `json:"lon"`
	Name    string  `json:"name"`
	Region  string  `json:"region"`
	Url     string  `json:"url"`
}

type TypeForecastday struct {
	Date string     `json:"date"`
	Day  TypeDay    `json:"day"`
	Hour []TypeHour `json:"hour"`
}
