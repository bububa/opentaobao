package drugtrace

import (
	"sync"
)

// AlibabaAlihealthDrugKytWesGetdruginfoDownloadurlResultModel 结构体
type AlibabaAlihealthDrugKytWesGetdruginfoDownloadurlResultModel struct {
	// 接口调用失败具体信息
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// 接口调用失败具体code
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// 返回值
	Model *JSONObject `json:"model,omitempty" xml:"model,omitempty"`
	// true：接口调用成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaAlihealthDrugKytWesGetdruginfoDownloadurlResultModel = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthDrugKytWesGetdruginfoDownloadurlResultModel)
	},
}

// GetAlibabaAlihealthDrugKytWesGetdruginfoDownloadurlResultModel() 从对象池中获取AlibabaAlihealthDrugKytWesGetdruginfoDownloadurlResultModel
func GetAlibabaAlihealthDrugKytWesGetdruginfoDownloadurlResultModel() *AlibabaAlihealthDrugKytWesGetdruginfoDownloadurlResultModel {
	return poolAlibabaAlihealthDrugKytWesGetdruginfoDownloadurlResultModel.Get().(*AlibabaAlihealthDrugKytWesGetdruginfoDownloadurlResultModel)
}

// ReleaseAlibabaAlihealthDrugKytWesGetdruginfoDownloadurlResultModel 释放AlibabaAlihealthDrugKytWesGetdruginfoDownloadurlResultModel
func ReleaseAlibabaAlihealthDrugKytWesGetdruginfoDownloadurlResultModel(v *AlibabaAlihealthDrugKytWesGetdruginfoDownloadurlResultModel) {
	v.MsgInfo = ""
	v.MsgCode = ""
	v.Model = nil
	v.Success = false
	poolAlibabaAlihealthDrugKytWesGetdruginfoDownloadurlResultModel.Put(v)
}
