/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"io"
	"strings"

	ixid "github.com/msaboor35/xidcli/internal/xid"
	"github.com/spf13/cobra"
)

func handleFlags(cmd *cobra.Command, xidArg string) error {
	time, err := cmd.Flags().GetBool("time")
	if err != nil {
		return err
	}
	machine, err := cmd.Flags().GetBool("machine")
	if err != nil {
		fmt.Println(err)
		return err
	}
	pid, err := cmd.Flags().GetBool("pid")
	if err != nil {
		fmt.Println(err)
		return err
	}
	counter, err := cmd.Flags().GetBool("counter")
	if err != nil {
		fmt.Println(err)
		return err
	}
	unixTime, err := cmd.Flags().GetBool("unix-time")
	if err != nil {
		fmt.Println(err)
		return err
	}

	if time {
		t, err := ixid.Time(xidArg)
		if err != nil {
			return err
		}
		fmt.Println(t.String())
	} else if machine {
		m, err := ixid.Machine(xidArg)
		if err != nil {
			return err
		}
		fmt.Println(string(m))
	} else if pid {
		p, err := ixid.Pid(xidArg)
		if err != nil {
			return err
		}
		fmt.Println(p)
	} else if counter {
		c, err := ixid.Counter(xidArg)
		if err != nil {
			return err
		}
		fmt.Println(c)
	} else if unixTime {
		u, err := ixid.UnixTime(xidArg)
		if err != nil {
			return err
		}
		fmt.Println(u)
	}

	return nil
}

// parseCmd represents the parse command
var parseCmd = &cobra.Command{
	Use:   "parse [xid]",
	Short: "Parse xid",
	Long: `Parse subfields of xid.
	
	With no xid, read standard input`,
	ValidArgs: []string{"xid"},
	Run: func(cmd *cobra.Command, args []string) {
		var xidArg string
		if len(args) == 0 {
			stdin := cmd.InOrStdin()
			xidArgsBytes, err := io.ReadAll(stdin)
			if err != nil {
				fmt.Println(err)
				return
			}

			xidArg = string(xidArgsBytes)
		} else {
			xidArg = args[0]
		}

		xidArg = strings.TrimSpace(xidArg)

		err := handleFlags(cmd, xidArg)
		if err != nil {
			fmt.Println(err)
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(parseCmd)

	parseCmd.Flags().BoolP("time", "t", false, "Timestamp of xid")
	parseCmd.Flags().BoolP("unix-time", "u", false, "Timestamp of xid")
	parseCmd.Flags().BoolP("machine", "m", false, "Machine ID of xid")
	parseCmd.Flags().BoolP("pid", "p", false, "PID of xid")
	parseCmd.Flags().BoolP("counter", "c", false, "Counter of xid")
	parseCmd.MarkFlagsMutuallyExclusive("time", "machine", "pid", "counter", "unix-time")
	parseCmd.MarkFlagsOneRequired("time", "machine", "pid", "counter", "unix-time")
}
