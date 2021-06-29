package product

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
product匹配接口 API返回值 
tmall.product.schema.match

根据tmall.product.match.schema.get获取到的规则，填充相应地的字段值以及类目，匹配符合条件的产品，返回匹配product结果，注意，有可能返回多个产品ID，以逗号分隔（尤其是图书类目）；
*/
type TmallProductSchemaMatchAPIResponse struct {
    model.CommonResponse
    TmallProductSchemaMatchResponse
}

// product匹配接口 成功返回结果
type TmallProductSchemaMatchResponse struct {
    XMLName xml.Name `xml:"tmall_product_schema_match_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 返回匹配产品ID，部分类目可能返回多个产品ID，以逗号分隔。
    MatchResult   string `json:"match_result,omitempty" xml:"match_result,omitempty"`
}
