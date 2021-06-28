package category

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
查询商家被授权品牌列表和类目列表 APIResponse
taobao.itemcats.authorize.get

查询B商家被授权品牌列表、类目列表和 c 商家新品类目列表
*/
type TaobaoItemcatsAuthorizeGetAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"itemcats_authorize_get_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 里面有3个数组：<br/>Brand[]品牌列表,<br/>ItemCat[] 类目列表<br/>XinpinItemCat[] 针对于C卖家新品类目列表
    
    SellerAuthorize   *SellerAuthorize `json:"seller_authorize,omitempty" xml:"