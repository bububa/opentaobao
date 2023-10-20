package tuanhotel

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tuanhotel"
)

// Alitriptuanhotelimageupload 图片上传接口
// alitrip.tuan.hotel.image.upload
//
// 用户调用此接口完成外网图片上传至卖家图片中心，此接口返回图片中心的图片地址
func Alitriptuanhotelimageupload(clt *core.SDKClient, req *tuanhotel.AlitriptuanhotelimageuploadAPIRequest, session string) (*tuanhotel.AlitriptuanhotelimageuploadAPIResponse, error) {
	var resp tuanhotel.AlitriptuanhotelimageuploadAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
