package baichuan

// AsoDeviceInfoDo 
type AsoDeviceInfoDo struct {
    // imei
    Imei   string `json:"imei,omitempty" xml:"imei,omitempty"`
    // imsi
    Imsi   string `json:"imsi,omitempty" xml:"imsi,omitempty"`
    // idfa
    Idfa   string `json:"idfa,omitempty" xml:"idfa,omitempty"`
}
