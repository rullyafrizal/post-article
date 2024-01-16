package http

import (
	"github.com/gofiber/fiber/v2"
	"net/http"
	"strconv"
	"time"
)

type response struct {
	Code       string   `json:"code"`
	Message    string   `json:"message"`
	Data       any      `json:"data"`
	Errors     []string `json:"errors"`
	ServerTime int64    `json:"serverTime"`
}

func ResponseSuccess(ctx *fiber.Ctx, data any) error {
	return ctx.Status(fiber.StatusOK).JSON(&response{
		Code:       "00",
		Message:    "SUCCESS",
		Data:       data,
		ServerTime: time.Now().Unix(),
	})
}

func ResponseBadRequest(ctx *fiber.Ctx, err []error) error {
	return Error(ctx, http.StatusBadRequest, err, "Bad Request")
}

func ResponseNotFound(ctx *fiber.Ctx, err []error) error {
	return Error(ctx, http.StatusNotFound, err, "Resource Not Found")
}

func ResponseConflict(ctx *fiber.Ctx, err []error) error {
	return Error(ctx, http.StatusConflict, err, "Conflict")
}

func ResponseInternalServerError(ctx *fiber.Ctx, err []error) error {
	return Error(ctx, http.StatusInternalServerError, err, "Internal Server Error")
}

func ResponseStatusUnprocessableEntityError(ctx *fiber.Ctx, err []error) error {
	return Error(ctx, http.StatusUnprocessableEntity, err, "Unprocessable Entity")
}

func Error(ctx *fiber.Ctx, httpCode int, err []error, msg string) error {
	return ctx.Status(httpCode).JSON(&response{
		Code:       strconv.Itoa(httpCode),
		Message:    msg,
		Errors:     errorMessages(err),
		ServerTime: time.Now().Unix(),
	})
}

func ResponseError(ctx *fiber.Ctx, err []error, errType ErrorType) error {
	switch errType {
	case TypeBadRequest:
		return ResponseBadRequest(ctx, err)
	case TypeNotFound:
		return ResponseNotFound(ctx, err)
	case TypeConflict:
		return ResponseConflict(ctx, err)
	case TypeInternalServerError:
		return ResponseInternalServerError(ctx, err)
	case TypeUnprocessableEntity:
		return ResponseStatusUnprocessableEntityError(ctx, err)
	default:
		return ResponseInternalServerError(ctx, err)
	}
}

func errorMessages(err []error) []string {
	var errMessages []string
	for _, e := range err {
		errMessages = append(errMessages, e.Error())
	}
	return errMessages
}
