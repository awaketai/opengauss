package opengauss

var textTypesMap = map[string]string{
	"TEXT":       "text",
	"TINYTEXT":   "text",
	"MEDIUMTEXT": "text",
	"LONGTEXT":   "text",
}

var dateTypesMap = map[string]string{
	"DATE":      "date",
	"TIME":      "time",
	"DATETIME":  "timestamp",
	"TIMESTAMP": "timestamp",
}

var numericTypesMap = map[string]string{
	"TINYINT":   "tinyint",
	"SMALLINT":  "smallint",
	"INT":       "integer",
	"INTEGER":   "integer",
	"MEDIUMINT": "integer",
	"BIGINT":    "bigint",
	"FLOAT":     "real",
	"DOUBLE":    "float8",
}
