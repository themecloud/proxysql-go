package proxysql

import (
	"reflect"
	"testing"
)

func TestTable(t *testing.T) {
	expected := "table"
	result := Table(expected)(defaultHostQuery()).table
	if result != expected {
		t.Fatalf("did not set table properly: %s", result)
	}
}

func TestHostgroup(t *testing.T) {
	expected := 1
	result := Hostgroup(expected)(defaultHostQuery()).host.hostgroup_id
	if result != expected {
		t.Fatalf("did not set hostgroup properly: %d", result)
	}
}

func TestPort(t *testing.T) {
	expected := 3305
	result := Port(expected)(defaultHostQuery()).host.port
	if result != expected {
		t.Fatalf("did not set port properly: %d", result)
	}
}

func TestBuildAndParseEmptyHostQuery(t *testing.T) {
	opts, err := buildAndParseHostQuery()
	if err != nil {
		t.Logf("unexpected err: %v", err)
		t.Fail()
	}
	if !reflect.DeepEqual(opts, defaultHostQuery()) {
		t.Fatalf("parsed opts were not default: %v", opts)
	}
}

func TestBuildAndParseHostQueryWithHostgroup(t *testing.T) {
	opts, err := buildAndParseHostQuery(Hostgroup(1))
	if err != nil {
		t.Logf("unexpected err: %v", err)
		t.Fail()
	}
	if !reflect.DeepEqual(opts, defaultHostQuery().Hostgroup(1)) {
		t.Fatalf("parsed opts were not default: %v", opts)
	}
}

func TestBuildAndParseHostQueryError(t *testing.T) {
	opts, err := buildAndParseHostQuery(Hostgroup(-1))
	if err != ErrConfigBadHostgroup {
		t.Logf("did not receive expected err: %v", err)
		t.Fail()
	}
	if opts != nil {
		t.Fatalf("returned query obj was not null: %v", opts)
	}
}

func TestBuildSpecifiedColumns(t *testing.T) {
	opts, err := buildAndParseHostQuery(Hostgroup(1), Port(12))
	if err != nil {
		t.Logf("unexpected parse error: %v", err)
		t.Fail()
	}

	queryString := buildSpecifiedColumns(opts.specifiedFields)
	expected := "(hostgroup_id, port)"
	if expected != queryString {
		t.Fatalf("specified fields returned were not expected: %s != %s", expected, queryString)
	}
}

func TestBuildSpecifiedColumnsDifferentOrder(t *testing.T) {
	opts, err := buildAndParseHostQuery(Port(1), Hostgroup(12))
	if err != nil {
		t.Logf("unexpected parse error: %v", err)
		t.Fail()
	}

	queryString := buildSpecifiedColumns(opts.specifiedFields)
	expected := "(port, hostgroup_id)"
	if expected != queryString {
		t.Fatalf("specified fields returned were not expected: %s != %s", expected, queryString)
	}
}

func TestBuildSpecifiedColumnsDoesntGiveUsDuplicates(t *testing.T) {
	opts, err := buildAndParseHostQuery(Port(1), Hostgroup(12), Port(2))
	if err != ErrConfigDuplicateSpec {
		t.Logf("unexpected parse error: %v", err)
		t.Fail()
	}

	if opts != nil {
		t.Fatalf("did not receive nil value: %v", opts)
	}
}

func TestBuildSpecifiedColumnsIsOrderDependent(t *testing.T) {
	opts, err := buildAndParseHostQuery(Port(1), Hostgroup(12))
	if err != nil {
		t.Logf("unexpected parse error: %v", err)
		t.Fail()
	}

	for i := 0; i < 10; i++ {
		queryString := buildSpecifiedColumns(opts.specifiedFields)
		expected := "(port, hostgroup_id)"
		if expected != queryString {
			t.Fatalf("specified fields returned were not expected: %s != %s", expected, queryString)
		}
	}
}
