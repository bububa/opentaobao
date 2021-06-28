package caipiao

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
录入参加送彩票店铺信息 APIResponse
taobao.caipiao.shop.info.input

录入参加送彩票店铺信息，如果录入成功，返回true，如果录入失败，返回false，后端会根据卖家id与赠送类型（sellerid_presenttype_uk）来决定是新增数据还是修改数据。
*/
type TaobaoCaipiaoShopInfoInputAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"caipiao_shop_info_input_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 录入操作是否成功
    
    InputResult   bool `json:"input_result,omitempty" xml:"