package jst

// CheckSellerChooseErpResponse 
type CheckSellerChooseErpResponse struct {

    // 状态结果</br> 您无权查询此商家信息<br> 已与当前服务商签约改地址服务</br>  当前商家改地址服务未签约服务商</br> 已与其他服务商签约改地址服务</br>
    Status   string `json:"status,omitempty"`

    // 商家绑定的appkey
    Appkey   string `json:"appkey,omitempty"`

    // app名称
    AppTitle   string `json:"app_title,omitempty"`

}
