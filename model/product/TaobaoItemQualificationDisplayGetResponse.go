package product

import (
    "github.com/bububa/opentaobao/model"
)

/* 
资质采集配置异步获取接口 APIResponse
taobao.item.qualification.display.get

根据类目，商品，属性等参与动态获得资质采集配置
*/
type TaobaoItemQualificationDisplayGetAPIResponse struct {
    model.CommonResponse
    Response *TaobaoItemQualificationDisplayGetResponse `json:"taobao_item_qualification_display_get_response,omitempty"`
}

type TaobaoItemQualificationDisplayGetResponse struct {

    // 返回资质采集配置
    DisplayConf   *DisplayQualifications `json:"display_conf,omitempty"`

}