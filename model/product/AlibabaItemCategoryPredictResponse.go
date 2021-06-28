package product

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
商品发布类目预测 APIResponse
alibaba.item.category.predict

<font color='red'>商品发布类目预测接口，预测匹配的结果存在一定误差，需要商家二次确认，避免类目配置错误产生其他影响。</font>
*/
type AlibabaItemCategoryPredictAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_item_category_predict_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 类目路径
    
    CatName   string `json:"cat_name,omitempty" xml:"