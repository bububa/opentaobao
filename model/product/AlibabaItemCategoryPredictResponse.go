package product

import (
    "github.com/bububa/opentaobao/model"
)

/* 
商品发布类目预测 APIResponse
alibaba.item.category.predict

<font color='red'>商品发布类目预测接口，预测匹配的结果存在一定误差，需要商家二次确认，避免类目配置错误产生其他影响。</font>
*/
type AlibabaItemCategoryPredictAPIResponse struct {
    model.CommonResponse
    Response *AlibabaItemCategoryPredictResponse `json:"alibaba_item_category_predict_response,omitempty"`
}

type AlibabaItemCategoryPredictResponse struct {

    // 类目路径
    CatName   string `json:"cat_name,omitempty"`

    // 类目ID
    CatId   int64 `json:"cat_id,omitempty"`

    // 类目名称
    CatPath   string `json:"cat_path,omitempty"`

}
