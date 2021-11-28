package googleauth

import (
	"context"

	"github.com/NpoolPlatform/verification-door/message/npool"
	usersecret "github.com/NpoolPlatform/verification-door/pkg/crud/user-secret"
	"github.com/NpoolPlatform/verification-door/pkg/google"
	"golang.org/x/xerrors"
)

func GetQRcodeURL(ctx context.Context, in *npool.GetQRcodeURLRequest) (*npool.GetQRcodeURLResponse, error) {
	secret, err := google.GetSecret()
	if err != nil {
		return nil, xerrors.Errorf("fail to generate secret: %v", err)
	}

	secret, err = usersecret.Create(ctx, secret, in.UserID, in.AppID)
	if err != nil && err.Error() != usersecret.SecretExistError {
		return nil, xerrors.Errorf("fail to add user secret: %v", err)
	}

	qrcodeURL := google.GetQrcodeURL(in.Username, secret)

	return &npool.GetQRcodeURLResponse{
		Info: qrcodeURL,
	}, nil
}

func VerifyGoogleAuth(ctx context.Context, in *npool.VerifyGoogleAuthRequest) (*npool.VerifyGoogleAuthResponse, error) {
	secret, err := usersecret.GetUserSecret(ctx, in.UserID, in.AppID)
	if err != nil {
		return nil, xerrors.Errorf("fail to get user secret")
	}

	if result, err := google.VerifyCode(secret, in.Code); err != nil || !result {
		return nil, xerrors.Errorf("fail to verify code: %v", err)
	}

	return &npool.VerifyGoogleAuthResponse{
		Info: "google authentication verify successfully",
	}, nil
}
