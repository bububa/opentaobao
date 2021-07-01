package scs

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaScsImageMatteAPIRequest
阿里妈妈智能创意平台在线抠图 API请求
alibaba.scs.image.matte

该API对外输出一个在线抠图(Deep Image Matting)接口，合作方可以通过该接口利用深度学习抠图算法从图片中抠出目标对象(比如商品或者人物轮廓) */
type AlibabaScsImageMatteAPIRequest struct {
	model.Params
	// 资源位ID，接入前由智能创意平台分配
	_pid string
	// 服务名称，可选值: scs
	_name string
	// 场景名称，可选值: image_cutout
	_scenes string
	// 抠图上下文信息，json字符串格式，json中matting_type字段可选值: external_matting，url: 需要抠图的目标图片url
	_obExt string
	// 32位uuid
	_sessionid string
	// 当前秒级时间戳
	_ts string
}

// New
