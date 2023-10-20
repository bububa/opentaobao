package drugtrace

import (
	"sync"
)

// AlibabaAlihealthDrugtraceTopLsydGetkeyflagdruginfoDownloadurlResultModel 结构体
type AlibabaAlihealthDrugtraceTopLsydGetkeyflagdruginfoDownloadurlResultModel struct {
	// 接口调用失败具体信息
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// 接口调用失败具体code
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// 返回值
	Model *JSONObject `json:"model,omitempty" xml:"model,omitempty"`
	// true：接口调用成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaAlihealthDrugtraceTopLsydGetkeyflagdruginfoDownloadurlResultModel = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthDrugtraceTopLsydGetkeyflagdruginfoDownloadurlResultModel)
	},
}

// GetAlibabaAlihealthDrugtraceTopLsydGetkeyflagdruginfoDownloadurlResultModel() 从对象池中获取AlibabaAlihealthDrugtraceTopLsydGetkeyflagdruginfoDownloadurlResultModel
func GetAlibabaAlihealthDrugtraceTopLsydGetkeyflagdruginfoDownloadurlResultModel() *AlibabaAlihealthDrugtraceTopLsydGetkeyflagdruginfoDownloadurlResultModel {
	return poolAlibabaAlihealthDrugtraceTopLsydGetkeyflagdruginfoDownloadurlResultModel.Get().(*AlibabaAlihealthDrugtraceTopLsydGetkeyflagdruginfoDownloadurlResultModel)
}

// ReleaseAlibabaAlihealthDrugtraceTopLsydGetkeyflagdruginfoDownloadurlResultModel 释放AlibabaAlihealthDrugtraceTopLsydGetkeyflagdruginfoDownloadurlResultModel
func ReleaseAlibabaAlihealthDrugtraceTopLsydGetkeyflagdruginfoDownloadurlResultModel(v *AlibabaAlihealthDrugtraceTopLsydGetkeyflagdruginfoDownloadurlResultModel) {
	v.MsgInfo = ""
	v.MsgCode = ""
	v.Model = nil
	v.Success = false
	poolAlibabaAlihealthDrugtraceTopLsydGetkeyflagdruginfoDownloadurlResultModel.Put(v)
}
