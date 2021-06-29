package tvupadmin

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取终端类型下品牌列表 APIResponse
yunos.tvpubadmin.device.brands

获取终端类型下品牌列表
*/
type YunosTvpubadminDeviceBrandsAPIResponse struct {
    model.CommonResponse
    YunosTvpubadminDeviceBrandsResponse
}

type YunosTvpubadminDeviceBrandsResponse struct {
    XMLName xml.Name `xml:"yunos_tvpubadmin_device_brands_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // object
    
    List   []BrandDO `json:"list,omitempty" xml:"list>brand_do,omitempty"`
    
    
}
