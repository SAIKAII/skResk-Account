package accounts

import (
	"errors"
	"sync"

	"github.com/shopspring/decimal"

	"github.com/SAIKAII/skResk-Account/services"
	"github.com/SAIKAII/skResk-Infra/base"
)

var _ services.AccountService = new(accountService)
var once sync.Once

func init() {
	once.Do(func() {
		services.IAccountService = &accountService{}
	})
}

type accountService struct {
}

func (a *accountService) CreateAccount(dto services.AccountCreatedDTO) (*services.AccountDTO, error) {
	domain := accountDomain{}
	// 验证输入的参数是否合法
	if err := base.ValidateStruct(&dto); err != nil {
		return nil, err
	}
	// 执行账户创建的业务逻辑
	amount, err := decimal.NewFromString(dto.Amount)
	if err != nil {
		return nil, err
	}
	account := services.AccountDTO{
		AccountCreatedDTO: services.AccountCreatedDTO{
			UserId:       dto.UserId,
			Username:     dto.Username,
			AccountType:  dto.AccountType,
			AccountName:  dto.AccountName,
			CurrencyCode: dto.CurrencyCode,
		},
		Balance: amount,
		Status:  1,
	}
	rdto, err := domain.Create(account)
	return rdto, err
}

func (a *accountService) Transfer(dto services.AccountTransferDTO) (services.TransferedStatus, error) {
	// 验证参数
	domain := accountDomain{}
	// 验证输入的参数是否合法
	if err := base.ValidateStruct(&dto); err != nil {
		return services.TransferedStatusFailure, err
	}
	// 执行转账逻辑
	amount, err := decimal.NewFromString(dto.AmountStr)
	if err != nil {
		return services.TransferedStatusFailure, err
	}
	dto.Amount = amount
	if dto.ChangeFlag == services.FlagTransferOut {
		if dto.ChangeType > 0 {
			return services.TransferedStatusFailure,
				errors.New("如果changeflag为支出，那么changetype必须小于0")
		}
	} else {
		if dto.ChangeType < 0 {
			return services.TransferedStatusFailure,
				errors.New("如果changeflag为收入，那么changetype必须大于0")
		}
	}

	status, err := domain.Transfer(dto)
	return status, err
}

func (a *accountService) StoreValue(dto services.AccountTransferDTO) (services.TransferedStatus, error) {
	dto.TradeTarget = dto.TradeBody
	dto.ChangeFlag = services.FlagTransferIn
	dto.ChangeType = services.AccountStoreValue
	return a.Transfer(dto)
}

func (a *accountService) GetEnvelopeAccountByUserId(userId string) *services.AccountDTO {
	domain := accountDomain{}
	account := domain.GetEnvelopeAccountByUserId(userId)
	return account
}

func (a *accountService) GetAccount(accountNo string) *services.AccountDTO {
	domain := accountDomain{}
	return domain.GetAccount(accountNo)
}
