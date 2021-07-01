package alilabs

// TopRefreshReqDto 
type TopRefreshReqDto struct {
    // clientId
    ClientId   string `json:"client_id,omitempty" xml:"client_id,omitempty"`
    // 只支持“basic”
    Scope   string `json:"scope,omitempty" xml:"scope,omitempty"`
    // 只支持“refresh_token”
    GrantType   string `json:"grant_type,omitempty" xml:"grant_type,omitempty"`
    // refreshToken
    RefreshToken   string `json:"refresh_token,omitempty" xml:"refresh_token,omitempty"`
}
