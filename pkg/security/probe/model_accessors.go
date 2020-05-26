// Code generated - DO NOT EDIT.

package probe

import (
	"github.com/pkg/errors"

	"github.com/DataDog/datadog-agent/pkg/security/secl/eval"
)

var (
	ErrFieldNotFound = errors.New("field not found")
)

func (m *Model) GetEvaluator(key string) (interface{}, error) {
	switch key {

	case "container.id":

		return &eval.StringEvaluator{
			Eval:      func(ctx *eval.Context) string { return m.event.Container.ID },
			DebugEval: func(ctx *eval.Context) string { return m.event.Container.ID },

			Field: key,
		}, nil

	case "event.retval":

		return &eval.IntEvaluator{
			Eval:      func(ctx *eval.Context) int { return int(m.event.Event.Retval) },
			DebugEval: func(ctx *eval.Context) int { return int(m.event.Event.Retval) },

			Field: key,
		}, nil

	case "event.type":

		return &eval.StringEvaluator{
			Eval:      func(ctx *eval.Context) string { return m.event.Event.ResolveType(m.event.resolvers) },
			DebugEval: func(ctx *eval.Context) string { return m.event.Event.ResolveType(m.event.resolvers) },

			Field: key,
		}, nil

	case "mkdir.filename":

		return &eval.StringEvaluator{
			Eval:      func(ctx *eval.Context) string { return m.event.Mkdir.ResolveInode(m.event.resolvers) },
			DebugEval: func(ctx *eval.Context) string { return m.event.Mkdir.ResolveInode(m.event.resolvers) },

			Field: key,
		}, nil

	case "mkdir.inode":

		return &eval.IntEvaluator{
			Eval:      func(ctx *eval.Context) int { return int(m.event.Mkdir.Inode) },
			DebugEval: func(ctx *eval.Context) int { return int(m.event.Mkdir.Inode) },

			Field: key,
		}, nil

	case "mkdir.mode":

		return &eval.IntEvaluator{
			Eval:      func(ctx *eval.Context) int { return int(m.event.Mkdir.Mode) },
			DebugEval: func(ctx *eval.Context) int { return int(m.event.Mkdir.Mode) },

			Field: key,
		}, nil

	case "open.filename":

		return &eval.StringEvaluator{
			Eval:      func(ctx *eval.Context) string { return m.event.Open.ResolveInode(m.event.resolvers) },
			DebugEval: func(ctx *eval.Context) string { return m.event.Open.ResolveInode(m.event.resolvers) },

			Field: key,
		}, nil

	case "open.flags":

		return &eval.IntEvaluator{
			Eval:      func(ctx *eval.Context) int { return int(m.event.Open.Flags) },
			DebugEval: func(ctx *eval.Context) int { return int(m.event.Open.Flags) },

			Field: key,
		}, nil

	case "open.inode":

		return &eval.IntEvaluator{
			Eval:      func(ctx *eval.Context) int { return int(m.event.Open.Inode) },
			DebugEval: func(ctx *eval.Context) int { return int(m.event.Open.Inode) },

			Field: key,
		}, nil

	case "open.mode":

		return &eval.IntEvaluator{
			Eval:      func(ctx *eval.Context) int { return int(m.event.Open.Mode) },
			DebugEval: func(ctx *eval.Context) int { return int(m.event.Open.Mode) },

			Field: key,
		}, nil

	case "process.gid":

		return &eval.IntEvaluator{
			Eval:      func(ctx *eval.Context) int { return int(m.event.Process.GID) },
			DebugEval: func(ctx *eval.Context) int { return int(m.event.Process.GID) },

			Field: key,
		}, nil

	case "process.name":

		return &eval.StringEvaluator{
			Eval:      func(ctx *eval.Context) string { return m.event.Process.HandleComm(m.event.resolvers) },
			DebugEval: func(ctx *eval.Context) string { return m.event.Process.HandleComm(m.event.resolvers) },

			Field: key,
		}, nil

	case "process.pid":

		return &eval.IntEvaluator{
			Eval:      func(ctx *eval.Context) int { return int(m.event.Process.Pid) },
			DebugEval: func(ctx *eval.Context) int { return int(m.event.Process.Pid) },

			Field: key,
		}, nil

	case "process.pidns":

		return &eval.IntEvaluator{
			Eval:      func(ctx *eval.Context) int { return int(m.event.Process.Pidns) },
			DebugEval: func(ctx *eval.Context) int { return int(m.event.Process.Pidns) },

			Field: key,
		}, nil

	case "process.tid":

		return &eval.IntEvaluator{
			Eval:      func(ctx *eval.Context) int { return int(m.event.Process.Tid) },
			DebugEval: func(ctx *eval.Context) int { return int(m.event.Process.Tid) },

			Field: key,
		}, nil

	case "process.tty_name":

		return &eval.StringEvaluator{
			Eval:      func(ctx *eval.Context) string { return m.event.Process.HandleTTY(m.event.resolvers) },
			DebugEval: func(ctx *eval.Context) string { return m.event.Process.HandleTTY(m.event.resolvers) },

			Field: key,
		}, nil

	case "process.uid":

		return &eval.IntEvaluator{
			Eval:      func(ctx *eval.Context) int { return int(m.event.Process.UID) },
			DebugEval: func(ctx *eval.Context) int { return int(m.event.Process.UID) },

			Field: key,
		}, nil

	case "rename.newfilename":

		return &eval.StringEvaluator{
			Eval:      func(ctx *eval.Context) string { return m.event.Rename.ResolveTargetInode(m.event.resolvers) },
			DebugEval: func(ctx *eval.Context) string { return m.event.Rename.ResolveTargetInode(m.event.resolvers) },

			Field: key,
		}, nil

	case "rename.newinode":

		return &eval.IntEvaluator{
			Eval:      func(ctx *eval.Context) int { return int(m.event.Rename.TargetInode) },
			DebugEval: func(ctx *eval.Context) int { return int(m.event.Rename.TargetInode) },

			Field: key,
		}, nil

	case "rename.oldfilename":

		return &eval.StringEvaluator{
			Eval:      func(ctx *eval.Context) string { return m.event.Rename.ResolveSrcInode(m.event.resolvers) },
			DebugEval: func(ctx *eval.Context) string { return m.event.Rename.ResolveSrcInode(m.event.resolvers) },

			Field: key,
		}, nil

	case "rename.oldinode":

		return &eval.IntEvaluator{
			Eval:      func(ctx *eval.Context) int { return int(m.event.Rename.SrcInode) },
			DebugEval: func(ctx *eval.Context) int { return int(m.event.Rename.SrcInode) },

			Field: key,
		}, nil

	case "rmdir.filename":

		return &eval.StringEvaluator{
			Eval:      func(ctx *eval.Context) string { return m.event.Rmdir.ResolveInode(m.event.resolvers) },
			DebugEval: func(ctx *eval.Context) string { return m.event.Rmdir.ResolveInode(m.event.resolvers) },

			Field: key,
		}, nil

	case "rmdir.inode":

		return &eval.IntEvaluator{
			Eval:      func(ctx *eval.Context) int { return int(m.event.Rmdir.Inode) },
			DebugEval: func(ctx *eval.Context) int { return int(m.event.Rmdir.Inode) },

			Field: key,
		}, nil

	case "unlink.filename":

		return &eval.StringEvaluator{
			Eval:      func(ctx *eval.Context) string { return m.event.Unlink.ResolveInode(m.event.resolvers) },
			DebugEval: func(ctx *eval.Context) string { return m.event.Unlink.ResolveInode(m.event.resolvers) },

			Field: key,
		}, nil

	case "unlink.inode":

		return &eval.IntEvaluator{
			Eval:      func(ctx *eval.Context) int { return int(m.event.Unlink.Inode) },
			DebugEval: func(ctx *eval.Context) int { return int(m.event.Unlink.Inode) },

			Field: key,
		}, nil

	}

	return nil, errors.Wrap(ErrFieldNotFound, key)
}

func (m *Model) GetTags(key string) ([]string, error) {
	switch key {

	case "container.id":
		return []string{}, nil

	case "event.retval":
		return []string{}, nil

	case "event.type":
		return []string{}, nil

	case "mkdir.filename":
		return []string{}, nil

	case "mkdir.inode":
		return []string{}, nil

	case "mkdir.mode":
		return []string{}, nil

	case "open.filename":
		return []string{}, nil

	case "open.flags":
		return []string{}, nil

	case "open.inode":
		return []string{}, nil

	case "open.mode":
		return []string{}, nil

	case "process.gid":
		return []string{}, nil

	case "process.name":
		return []string{}, nil

	case "process.pid":
		return []string{}, nil

	case "process.pidns":
		return []string{}, nil

	case "process.tid":
		return []string{}, nil

	case "process.tty_name":
		return []string{}, nil

	case "process.uid":
		return []string{}, nil

	case "rename.newfilename":
		return []string{}, nil

	case "rename.newinode":
		return []string{}, nil

	case "rename.oldfilename":
		return []string{}, nil

	case "rename.oldinode":
		return []string{}, nil

	case "rmdir.filename":
		return []string{}, nil

	case "rmdir.inode":
		return []string{}, nil

	case "unlink.filename":
		return []string{}, nil

	case "unlink.inode":
		return []string{}, nil

	}

	return nil, errors.Wrap(ErrFieldNotFound, key)
}

func (m *Model) GetEventType(key string) (string, error) {
	switch key {

	case "container.id":
		return "container", nil

	case "event.retval":
		return "", nil

	case "event.type":
		return "", nil

	case "mkdir.filename":
		return "mkdir", nil

	case "mkdir.inode":
		return "mkdir", nil

	case "mkdir.mode":
		return "mkdir", nil

	case "open.filename":
		return "open", nil

	case "open.flags":
		return "open", nil

	case "open.inode":
		return "open", nil

	case "open.mode":
		return "open", nil

	case "process.gid":
		return "", nil

	case "process.name":
		return "", nil

	case "process.pid":
		return "", nil

	case "process.pidns":
		return "", nil

	case "process.tid":
		return "", nil

	case "process.tty_name":
		return "", nil

	case "process.uid":
		return "", nil

	case "rename.newfilename":
		return "rename", nil

	case "rename.newinode":
		return "rename", nil

	case "rename.oldfilename":
		return "rename", nil

	case "rename.oldinode":
		return "rename", nil

	case "rmdir.filename":
		return "rmdir", nil

	case "rmdir.inode":
		return "rmdir", nil

	case "unlink.filename":
		return "unlink", nil

	case "unlink.inode":
		return "unlink", nil

	}

	return "", errors.Wrap(ErrFieldNotFound, key)
}