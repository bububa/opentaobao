package product

import (
    "github.com/bububa/opentaobao/model"
)

/* 
导入数据到商品百科服务 APIResponse
taobao.baike.import.zhubao.data

用于接入外部数据录入到商品百科中
*/
type TaobaoBaikeImportZhubaoDataAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoBaikeImportZhubaoDataResponse `json:"baike_import_zhubao_data_response,omitempty"` 
    TaobaoBaikeImportZhubaoDataResponse
}

/* model for simplify = false
type TaobaoBaikeImportZhubaoDataResponse struct {

    // result
    
    Result  *struct {
        TaobaoBaikeImportZhubaoDataResult  *TaobaoBaikeImportZhubaoDataResult `json:"taobao_baike_import_zhubao_data_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type TaobaoBaikeImportZhubaoDataResponse struct {

    // result
    Result   *TaobaoBaikeImportZhubaoDataResult `json:"result,omitempty"`

}
