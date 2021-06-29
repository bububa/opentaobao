package alilabs

// DeviceTokenVo 
type DeviceTokenVo struct {

    // accessToken超时时间，单位为秒
    
    AccessExpiresIn   int64 `json:"access_expires_in,omitempty" xml:"access_expires_in,omitempty"`
    

    // accessToken
    
    AccessToken   string `json:"access_token,omitempty" xml:"access_token,omitempty"`
    

    // refreshToken超时时间，单位为秒
    
    RefreshExpiresIn   int64 `json:"refresh_expires_in,omitempty" xml:"refresh_expires_in,omitempty"`
    

    // refreshToken
    
    RefreshToken   string `json:"refresh_token,omitempty" xml:"refresh_token,omitempty"`
    

    // state
    
    State   string `json:"state,omitempty" xml:"state,omitempty"`
    

}
