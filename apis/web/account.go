package web

import (
	"github.com/SAIKAII/skResk-Account/services"
	"github.com/SAIKAII/skResk-Infra"
	"github.com/SAIKAII/skResk-Infra/base"
	"github.com/kataras/iris"
)

// 定义Web api的时候，对每一个子业务，定义统一的前缀
// 资金账户的根路径定义为：/account
// 版本号：/v1/account

const (
	ResCodeBizTransferedFailure = base.ResCodeBizError + 1
)

func init() {
	infra.RegisterApi(new(AccountApi))
}

type AccountApi struct {
	service services.AccountService
}

func (a *AccountApi) Init() {
	a.service = services.GetAccountService()
	groupRouter := base.Iris().Party("/v1/account")
	groupRouter.Post("/create", a.createHandler)
	groupRouter.Post("/transfer", a.transferHandler)
	groupRouter.Get("/envelope/get", a.getEnvelopeAccountHandler)
	groupRouter.Get("/get", a.getAccountHandler)
}

// 账户创建的接口：/v1/account/create
// POST body json
func (a *AccountApi) createHandler(ctx iris.Context) {
	// 获取请求参数
	account := services.AccountCreatedDTO{}
	err := ctx.ReadJSON(&account)
	r := base.Res{
		Code: base.ResCodeOk,
	}
	if err != nil {
		r.Code = base.ResCodeRequestParamsError
		r.Message = err.Error()
		ctx.JSON(&r)
		return
	}
	// 执行创建账户的代码
	service := services.GetAccountService()
	dto, err := service.CreateAccount(account)
	if err != nil {
		r.Code = base.ResCodeInnerServerError
		r.Message = err.Error()
	}
	r.Data = dto
	ctx.JSON(&r)

}

// 转账的接口：/v1/account/transfer
// POST
func (a *AccountApi) transferHandler(ctx iris.Context) {
	// 获取请求参数
	transfer := services.AccountTransferDTO{}
	err := ctx.ReadJSON(&transfer)
	r := base.Res{
		Code: base.ResCodeOk,
	}
	if err != nil {
		r.Code = base.ResCodeRequestParamsError
		r.Message = err.Error()
		ctx.JSON(&r)
		return
	}
	// 执行转账逻辑
	service := services.GetAccountService()
	status, err := service.Transfer(transfer)
	if err != nil {
		r.Code = base.ResCodeInnerServerError
		r.Message = err.Error()
	}
	r.Data = status
	if status != services.TransferedStatusSuccess {
		r.Code = ResCodeBizTransferedFailure
		r.Message = err.Error()
	}
	ctx.JSON(&r)
}

// 查询红包账户的接口：/v1/account/envelope/get
func (a *AccountApi) getEnvelopeAccountHandler(ctx iris.Context) {
	userId := ctx.URLParam("userId")
	r := base.Res{
		Code: base.ResCodeOk,
	}
	if userId == "" {
		r.Code = base.ResCodeRequestParamsError
		r.Message = "用户ID"
		ctx.JSON(&r)
		return
	}
	service := services.GetAccountService()
	account := service.GetEnvelopeAccountByUserId(userId)
	if account == nil {
		r.Code = base.ResCodeInnerServerError
		r.Message = "无法获取指定用户"
	} else {
		r.Data = account
	}
	ctx.JSON(&r)
}

// 查询账户信息的接口：/v1/account/get
func (a *AccountApi) getAccountHandler(ctx iris.Context) {
	// 获取请求参数
	accountNo := ctx.URLParam("accountNo")
	r := base.Res{
		Code: base.ResCodeOk,
	}
	if accountNo == "" {
		r.Code = base.ResCodeRequestParamsError
		r.Message = "账户编号不能为空"
		ctx.JSON(&r)
		return
	}
	service := services.GetAccountService()
	account := service.GetAccount(accountNo)
	if account == nil {
		r.Code = base.ResCodeInnerServerError
		r.Message = "无法获取指定账户"
	} else {
		r.Data = account
	}
	ctx.JSON(&r)
}
