package fenxiao

import (
    "github.com/bububa/opentaobao/model"
)

/* 
创建IC商品与后端商品的映射关系 APIResponse
taobao.scitem.map.add

创建IC商品或分销商品与后端商品的映射关系
*/
type TaobaoScitemMapAddAPIResponse struct {
    model.CommonResponse
    Response *TaobaoScitemMapAddResponse `json:"taobao_scitem_map_add_response,omitempty"`
}

type TaobaoScitemMapAddResponse struct {

    // 接口调用返回结果信息：商家编码
    OuterCode   string `json:"outer_code,omitempty"`

}
