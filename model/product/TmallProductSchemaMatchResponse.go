package product

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
product匹配接口 APIResponse
tmall.product.schema.match

根据tmall.product.match.schema.get获取到的规则，填充相应地的字段值以及类目，匹配符合条件的产品，返回匹配product结果，注意，有可能返回多个产品ID，以逗号分隔（尤其是图书类目）；
*/
type TmallProductSchemaMatchAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"tmall_product_schema_match_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 返回匹配产品ID，部分类目可能返回多个产品ID，以逗号分隔。
    
    MatchResult   string `json:"match_result,omitempty" xml:"