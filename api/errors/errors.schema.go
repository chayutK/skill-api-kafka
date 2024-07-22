package errs

type ErrorResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func NotFoundError() ErrorResponse {
	return ErrorResponse{
		Status:  "error",
		Message: "Skill not found",
	}
}

func InternalServerError(err error) ErrorResponse {
	return ErrorResponse{
		Status:  "error",
		Message: err.Error(),
	}
}

func AlreadyExistError() ErrorResponse {
	return ErrorResponse{
		Status:  "error",
		Message: "Skill already exists",
	}
}

func UpdatePutError() ErrorResponse {
	return ErrorResponse{
		Status:  "error",
		Message: "not be able to update skill",
	}
}

func UpdatePatchError(t string) ErrorResponse {
	return ErrorResponse{
		Status:  "error",
		Message: "not be able to update skill" + t,
	}
}

func BadRequestError() ErrorResponse {
	return ErrorResponse{
		Status:  "error",
		Message: "Your input is invalid",
	}
}
