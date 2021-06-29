package travel

// ItemRefundInfo 
type ItemRefundInfo struct {
    // 退改规则 1）格式：标准规则 或 自定义规则：a_a_num,a_b_num,b-1_c_num,c-1_c-1_num 2）规则：自定义规则里最多可含5组规则
    RefundRegulations   []string `json:"refund_regulations,omitempty" xml:"refund_regulations>string,omitempty"`
    // 退改规则类型，0为标准，1为自定义 2为不支持退改规则。不传默认为0
    RefundType   int64 `json:"refund_type,omitempty" xml:"refund_type,omitempty"`
    // 退款规则（json数组格式）。自定义退改时需填写（与refund_regulations字段二选一）。示例中一共包含4条规则（3条平日规则，1条节假日规则），按照顺序每条规则含义如下：出行前5日及以上，买家违约收取总费用的50，卖家违约收取总费用的20；出行前4日至1日，买家违约收取总费用的80，卖家违约收取总费用的50；行程开始当天，买家违约收取总费用的100，卖家违约收取总费用的70；如果行程日期包含节假日，则节假日条款为买家违约收取总费用的100，卖家违约收取总费用的90
    RefundRegulationsJson   string `json:"refund_regulations_json,omitempty" xml:"refund_regulations_json,omitempty"`
}
