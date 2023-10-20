package koubeimall

import (
	"sync"
)

// ItemServices 结构体
type ItemServices struct {
	// 服务名称列表
	ServiceNameList []string `json:"service_name_list,omitempty" xml:"service_name_list>string,omitempty"`
	// 服务明细，点击弹层显示
	ServicesDetailList []ItemServicesDetail `json:"services_detail_list,omitempty" xml:"services_detail_list>item_services_detail,omitempty"`
	// 服务列表
	ServiceCodeList []string `json:"service_code_list,omitempty" xml:"service_code_list>string,omitempty"`
}

var poolItemServices = sync.Pool{
	New: func() any {
		return new(ItemServices)
	},
}

// GetItemServices() 从对象池中获取ItemServices
func GetItemServices() *ItemServices {
	return poolItemServices.Get().(*ItemServices)
}

// ReleaseItemServices 释放ItemServices
func ReleaseItemServices(v *ItemServices) {
	v.ServiceNameList = v.ServiceNameList[:0]
	v.ServicesDetailList = v.ServicesDetailList[:0]
	v.ServiceCodeList = v.ServiceCodeList[:0]
	poolItemServices.Put(v)
}
