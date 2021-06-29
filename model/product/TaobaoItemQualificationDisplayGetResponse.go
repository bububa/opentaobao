package product

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
资质采集配置异步获取接口 APIResponse
taobao.item.qualification.display.get

根据类目，商品，属性等参与动态获得资质采集配置
*/
type TaobaoItemQualificationDisplayGetAPIResponse struct {
    model.CommonResponse
    TaobaoItemQualificationDisplayGetResponse
}

type TaobaoItemQualificationDisplayGetResponse struct {
    XMLName xml.Name `xml:"item_qualification_display_get_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 返回资质采集配置
    
    DisplayConf   *DisplayQualifications `json:"display_conf,omitempty" xml:"display_conf,omitempty"`

    
}
