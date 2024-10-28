package opengauss

import "testing"

func Test_textTypesMap_compatibilityMysqlSyntax(t *testing.T) {
	dl := new(Dialector)
	tests := []struct {
		name     string
		sqlType  string
		expected string
		ok       bool
	}{
		{"TINYTEXT", "TINYTEXT", "text", true},
		{"TINYTEXT NOT NULL", "TINYTEXT NOT NULL", "text", true},
		{"TEXT not null", "TEXT not null", "text", true},
		{"MEDIUMTEXT", "MEDIUMTEXT", "text", true},
		{"MEDIUMTEXT", "MEDIUMTEXT NOT NULL ", "text", true},
		{"MEDIUMTEXT", "MEDIUMTEXT NOT NULL DEFAULT 0", "text", true},
		{"LONGTEXT", "LONGTEXT", "text", true},
		{"LONGTEXT not null", "LONGTEXT", "text", true},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual, ok := dl.compatibilityMysqlSyntax(textTypesMap, test.sqlType)
			if ok != test.ok {
				t.Errorf("expected ok to be %v, got %v", test.ok, ok)
			}
			if actual != test.expected {
				t.Errorf("expected %v, got %v", test.expected, actual)
			}
		})
	}

}

func Test_dateTypesMap_compatibilityMysqlSyntax(t *testing.T) {
	dl := new(Dialector)
	tests := []struct {
		name     string
		sqlType  string
		expected string
		ok       bool
	}{
		{"DATE", "DATE", "date", true},
		{"DATE not null", "DATE not null", "date", true},
		{"DATE NOT null", "DATE NOT null", "date", true},
		{"DATE NOT NULL", "DATE NOT NULL", "date", true},
		{"TIME", "TIME", "time", true},
		{"TIME not null", "TIME not null", "time", true},
		{"DATETIME", "DATETIME", "timestamp", true},
		{"DATETIME not null", "DATETIME not null", "timestamp", true},
		{"TIMESTAMP", "TIMESTAMP", "timestamp", true},
		{"TIMESTAMP NOT NULL", "TIMESTAMP NOT NULL ", "timestamp", true},
		{"TIMESTAMP NOT NULL", "TIMESTAMP NOT NULL DEFAULT 0", "timestamp", true},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual, ok := dl.compatibilityMysqlSyntax(dateTypesMap, test.sqlType)
			if ok != test.ok {
				t.Errorf("expected ok to be %v, got %v", test.ok, ok)
			}
			if actual != test.expected {
				t.Errorf("expected %v, got %v", test.expected, actual)
			}
		})
	}
}

func Test_numericTypesMap_compatibilityMysqlSyntax(t *testing.T) {
	dl := new(Dialector)
	tests := []struct {
		name     string
		sqlType  string
		expected string
		ok       bool
	}{
		{"TINYINT", "TINYINT", "tinyint", true},
		{"TINYINT not null", "TINYINT not null", "tinyint", true},
		{"TINYINT NOT NULL null", "TINYINT NOT NULL", "tinyint", true},
		{"SMALLINT", "SMALLINT", "smallint", true},
		{"SMALLINT not null", "SMALLINT not null", "smallint", true},
		{"SMALLINT not null", "SMALLINT not null default 0", "smallint", true},
		{"INT ", "INT", "integer", true},
		{"INTEGER", "INTEGER", "integer", true},
		{"MEDIUMINT", "MEDIUMINT", "integer", true},
		{"MEDIUMINT NOT NULL", "MEDIUMINT NOT NULL", "integer", true},
		{"BIGINT", "BIGINT", "bigint", true},
		{"FLOAT", "FLOAT", "real", true},
		{"DOUBLE", "DOUBLE", "float8", true},
		{"UNSUPPORTED", "UNSUPPORTED", "UNSUPPORTED", false},
		{"VARCHAR", "VARCHAR(255)", "VARCHAR(255)", false},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual, ok := dl.compatibilityMysqlSyntax(numericTypesMap, test.sqlType)
			if ok != test.ok {
				t.Errorf("expected ok to be %v, got %v", test.ok, ok)
			}
			if actual != test.expected {
				t.Errorf("expected %v, got %v", test.expected, actual)
			}
		})
	}
}
