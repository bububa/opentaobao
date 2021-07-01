package product

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaItemCategoryPredictAPIResponse
商品发布类目预测 API返回值
alibaba.item.category.predict

<font color='red'>商品发布类目预测接口，预测匹配的结果存在一定误差，需要商家二次确认，避免类目配置错误产生其他影响。</font> */
type AlibabaItemCategoryPredictAPIResponse struct {
	model.CommonResponse
	AlibabaItemCategoryPredictAPIResponseModel
}

// AlibabaItemCategoryPredictAPIResponseModel is 商品发布类目预测 成功返回结果
type AlibabaItemCategoryPredictAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_item_category_predict_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 类目路径
	CatName string `json:"cat_name,omitempty" xml:"cat_name,omitempty"`
	// 类目ID
	CatId int64 `json:"cat_id,omitempty" xml:"cat_id,omitempty"`
	// 类目名称
	CatPath string `json:"cat_path,omitempty" xml:"cat_path,omitempty"`
}
