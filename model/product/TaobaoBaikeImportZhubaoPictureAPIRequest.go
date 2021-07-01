package product

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoBaikeImportZhubaoPictureAPIRequest
百科图片数据导入 API请求
taobao.baike.import.zhubao.picture

用于接入外部--图片--录入到商品百科中 */
type TaobaoBaikeImportZhubaoPictureAPIRequest struct {
	model.Params
	// 图片二进制数据
	_picture *model.File
}

// New
