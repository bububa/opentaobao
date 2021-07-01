package normalvisa

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoAlitripTravelNormalvisaUploadfileAPIRequest
上传电子签证 API请求
taobao.alitrip.travel.normalvisa.uploadfile

上传电子签证 */
type TaobaoAlitripTravelNormalvisaUploadfileAPIRequest struct {
	model.Params
	// 文件
	_fileBytes *model.File
	// 文件名：注意文件名请保证和上传的文件一直
	_fileName string
	// 订单id
	_bizOrderId int64
}

// New
