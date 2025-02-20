package messages

var EsMessages = map[MessageKeysEnum]string{
	"SOMETHING_WENT_WRONG":           "Oh, lo sentimos, algo no salió bien con esta acción!",
	"RESOURCE_NOT_FOUND":             "Recurso no encontrado!",
	"SOME_PARAMETERS_ARE_MISSING":    "Faltan algunos parámetros: %s.",
	"UNKNOWN_RESPONSE_STATUS":        "Estado de respuesta del servidor desconocido.",
	"TOOL_HAS_NOT_BEEN_INITIALIZED":  "La herramienta %s no ha sido configurada.",
	"PROCESSING_DATA_CLIENT_ERROR":   "Error procesando la data en el cliente http.",
	"DEPENDENCY_NOT_FOUND":           "No se encontró %s en el contenedor de dependencias.",
	"AUTHORIZATION_REQUIRED":         "La autorización es requerida.",
	"INVALID_USER_OR_PASSWORD":       "Usuario o contraseña incorrecto.",
	"ERROR_CREATING_USER":            "Algo no salió bien registrado el usuario.",
	"USER_WAS_CREATED":               "El usuario fue registrado",
	"USER_WITH_EMAIL_ALREADY_EXISTS": "Un usuario ya ha sido registrado con el email: %s.",
	"INVALID_EMAIL":                  "El correo electrónico es inválido.",
	"INVALID_PASSWORD":               "La contraseña es inválida.",
	"INVALID_SESSION":                "La sesión es inválida.",
	"APPLICATION_STATUS_OK":          "La aplicación está en ejecución.",
}
