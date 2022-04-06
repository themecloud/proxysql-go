package proxysql

import "fmt"

// Replication represents a row in ProxySQL's mysql_replication_hostgroups config table

type Replication struct {
	writer_hostgroup int
	reader_hostgroup int
	comment          string
}

// DefaultReplication returns a default host (in terms of the mysql_replication_hostgroups table).
// Note that comment

func DefaultReplication() *Replication {
	return &Replication{
		writer_hostgroup: 10,
		reader_hostgroup: 20,
		comment:          "",
	}
}

// Setters for Replication struct

func (r *Replication) SetWriterHostGroup(wh int) *Replication {
	r.writer_hostgroup = wh
	return r
}

func (r *Replication) SetReaderHostGroup(rh int) *Replication {
	r.reader_hostgroup = rh
	return r
}

func (r *Replication) SetComment(c string) *Replication {
	r.comment = c
	return r
}

// Builders for Host struct

func (r *Replication) WriterHostGroup() int {
	return r.writer_hostgroup
}

func (r *Replication) ReaderHostGroup() int {
	return r.reader_hostgroup
}

func (r *Replication) Comment() string {
	return r.comment
}

func (r *Replication) values() string {
	return fmt.Sprintf("(%d, %d, '%s')", r.writer_hostgroup, r.reader_hostgroup, r.comment)
}

func (r *Replication) columns() string {
	return "(writer_hostgroup, reader_hostgroup, comment)"
}

func (r *Replication) where() string {
	return fmt.Sprintf("writer_hostgroup = %d and reader_hostgroup = %d and comment = '%s'", r.writer_hostgroup, r.reader_hostgroup, r.comment)
}
