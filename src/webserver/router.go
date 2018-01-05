package webserver

import (
	"github.com/asepnur/meiko_user/src/util/auth"
	"github.com/asepnur/meiko_user/src/webserver/handler/rolegroup"
	"github.com/asepnur/meiko_user/src/webserver/handler/user"
	"github.com/julienschmidt/httprouter"
)

// Load returns all routing of this server
func loadRouter(r *httprouter.Router) {

	// ========================== User Handler ==========================
	// User section

	r.POST("/api/v1/user/register", auth.OptionalAuthorize(user.SignUpHandler))
	r.POST("/api/v1/user/verify", auth.OptionalAuthorize(user.EmailVerificationHandler))
	r.POST("/api/v1/user/signin", auth.OptionalAuthorize(user.SignInHandler))
	r.POST("/api/v1/user/forgot", auth.OptionalAuthorize(user.ForgotHandler))
	r.POST("/api/v1/user/signout", auth.MustAuthorize(user.SignOutHandler)) // delete
	r.POST("/api/v1/user/profile", auth.MustAuthorize(user.UpdateProfileHandler))
	r.GET("/api/v1/user/profile", auth.MustAuthorize(user.GetProfileHandler))
	r.POST("/api/v1/user/changepassword", auth.MustAuthorize(user.ChangePasswordHandler))
	// Admin section
	r.GET("/api/admin/v1/user", auth.MustAuthorize(user.ReadHandler))
	r.POST("/api/admin/v1/user", auth.MustAuthorize(user.CreateHandler))
	r.GET("/api/admin/v1/user/:id", auth.MustAuthorize(user.DetailHandler))
	r.PATCH("/api/admin/v1/user/:id", auth.MustAuthorize(user.UpdateHandler))
	r.PATCH("/api/admin/v1/user/:id/:status", auth.MustAuthorize(user.ActivationHandler))
	r.DELETE("/api/admin/v1/user/:id", auth.MustAuthorize(user.DeleteHandler))

	// Public
	r.GET("/api/v1/util/time", user.GetTimeHandler)
	// ======================== End User Handler ========================

	// ======================== Rolegroup Handler =======================
	// User Section
	r.GET("/api/v1/role", auth.OptionalAuthorize(rolegroup.GetPrivilege))
	// Admin section
	r.GET("/api/admin/v1/role", auth.MustAuthorize(rolegroup.ReadHandler))
	r.POST("/api/admin/v1/role", auth.MustAuthorize(rolegroup.CreateHandler))
	r.GET("/api/admin/v1/role/:rolegroup_id", auth.MustAuthorize(rolegroup.ReadDetailHandler))
	r.PATCH("/api/admin/v1/role/:rolegroup_id", auth.MustAuthorize(rolegroup.UpdateHandler))
	r.DELETE("/api/admin/v1/role/:rolegroup_id", auth.MustAuthorize(rolegroup.DeleteHandler))
	// ====================== End Rolegroup Handler =====================

	// ========================== User Handler Internal ==========================
	// user internal

	r.POST("/api/v1/user/exhange-profile", auth.Oauth(user.ExchangeProfile))
	r.POST("/api/v1/user/exhange-id", auth.Oauth(user.ExchangeUserByID))
	r.POST("/api/v1/user/schedule-id", auth.Oauth(user.ExchangeUserByScheduleID))
	r.POST("/api/v1/user/identity-code", auth.Oauth(user.ExchangeUserByIdentityCode))
	// ====================== End  =====================
}
