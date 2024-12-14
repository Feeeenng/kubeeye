package output

import (
	"fmt"
	kubeeyev1alpha2 "github.com/kubesphere/kubeeye/apis/kubeeye/v1alpha2"
	"github.com/kubesphere/kubeeye/pkg/constant"
	"github.com/xuri/excelize/v2"
	corev1 "k8s.io/api/core/v1"
	"path"
)

func GenerateExcel(resultData *kubeeyev1alpha2.InspectResult, nodes *corev1.NodeList, pods *corev1.PodList) error {
	f := excelize.NewFile()
	defer func() error {
		if err := f.Close(); err != nil {
			return err
		}
		return nil
	}()
	// create a new sheet for Nodes
	if len(nodes.Items) > 0 {
		index, err := f.NewSheet(constant.NodesStatus)
		if err != nil {
			return err
		}

		nodeStatus := GetNodesStatus(nodes)

		// Write the data to the sheet
		for i, item := range nodeStatus {
			if i == 0 {
				for j, c := range item.Children {
					f.SetCellValue(constant.NodesStatus, fmt.Sprintf("%c1", 'A'+rune(j)), c.Text)
				}
			} else {
				for j, c := range item.Children {
					f.SetCellValue(constant.NodesStatus, fmt.Sprintf("%c%d", 'A'+rune(j), i+1), c.Text)
				}
			}
		}
		f.SetActiveSheet(index)
	}

	// Create a new sheet for Pods
	if len(pods.Items) > 0 {
		_, err := f.NewSheet(constant.AbnormalPods)
		if err != nil {
			return err
		}

		podStatus := GetAbnormalPods(pods)

		// Write the data to the sheet
		for i, item := range podStatus {
			if i == 0 {
				for j, c := range item.Children {
					f.SetCellValue(constant.AbnormalPods, fmt.Sprintf("%c1", 'A'+rune(j)), c.Text)
				}
			} else {
				for j, c := range item.Children {
					f.SetCellValue(constant.AbnormalPods, fmt.Sprintf("%c%d", 'A'+rune(j), i+1), c.Text)
				}
			}
		}
	}

	// Create a new sheet for OPA
	if resultData.Spec.OpaResult.ResourceResults != nil {
		_, err := f.NewSheet(constant.Opa)
		if err != nil {
			return err
		}

		opaResults := GetOpaList(resultData.Spec.OpaResult.ResourceResults)

		// Write the data to the sheet
		for i, item := range opaResults {
			if i == 0 {
				for j, c := range item.Children {
					f.SetCellValue(constant.Opa, fmt.Sprintf("%c1", 'A'+rune(j)), c.Text)
				}
			} else {
				for j, c := range item.Children {
					f.SetCellValue(constant.Opa, fmt.Sprintf("%c%d", 'A'+rune(j), i+1), c.Text)
				}
			}
		}
	}

	// Create a new sheet for systemd
	if resultData.Spec.SystemdResult != nil {
		_, err := f.NewSheet(constant.Systemd)
		if err != nil {
			return err
		}

		systemdResults := GetSystemd(resultData.Spec.SystemdResult)

		// Write the data to the sheet
		for i, item := range systemdResults {
			if i == 0 {
				for j, c := range item.Children {
					f.SetCellValue(constant.Systemd, fmt.Sprintf("%c1", 'A'+rune(j)), c.Text)
				}
			} else {
				for j, c := range item.Children {
					f.SetCellValue(constant.Systemd, fmt.Sprintf("%c%d", 'A'+rune(j), i+1), c.Text)
				}
			}
		}
	}

	// Create a new sheet for prometheus
	if resultData.Spec.PrometheusResult != nil {
		_, err := f.NewSheet(constant.Prometheus)
		if err != nil {
			return err
		}

		prometheusReults := GetPrometheus(resultData.Spec.PrometheusResult)
		// Write the data to the sheet
		for i, item := range prometheusReults {
			if i == 0 {
				for j, c := range item.Children {
					f.SetCellValue(constant.Prometheus, fmt.Sprintf("%c1", 'A'+rune(j)), c.Text)
				}
			} else {
				for j, c := range item.Children {
					f.SetCellValue(constant.Prometheus, fmt.Sprintf("%c%d", 'A'+rune(j), i+1), c.Text)
				}
			}
		}
	}

	// Create a new sheet for fileChange
	if resultData.Spec.FileChangeResult != nil {
		_, err := f.NewSheet(constant.FileChange)
		if err != nil {
			return err
		}

		fileChangeResults := GetFileChange(resultData.Spec.FileChangeResult)
		// Write the data to the sheet
		for i, item := range fileChangeResults {
			if i == 0 {
				for j, c := range item.Children {
					f.SetCellValue(constant.FileChange, fmt.Sprintf("%c1", 'A'+rune(j)), c.Text)
				}
			} else {
				for j, c := range item.Children {
					f.SetCellValue(constant.FileChange, fmt.Sprintf("%c%d", 'A'+rune(j), i+1), c.Text)
				}
			}
		}
	}

	// Create a new sheet for sysctl
	if resultData.Spec.SysctlResult != nil {
		_, err := f.NewSheet(constant.Sysctl)
		if err != nil {
			return err
		}

		sysctlResults := GetSysctl(resultData.Spec.SysctlResult)
		// Write the data to the sheet
		for i, item := range sysctlResults {
			if i == 0 {
				for j, c := range item.Children {
					f.SetCellValue(constant.Sysctl, fmt.Sprintf("%c1", 'A'+rune(j)), c.Text)
				}
			} else {
				for j, c := range item.Children {
					f.SetCellValue(constant.Sysctl, fmt.Sprintf("%c%d", 'A'+rune(j), i+1), c.Text)
				}
			}
		}
	}

	// Create a new sheet for filefilter
	if resultData.Spec.FileFilterResult != nil {
		_, err := f.NewSheet(constant.FileFilter)
		if err != nil {
			return err
		}

		fileFilterResults := GetFileFilter(resultData.Spec.FileFilterResult)
		// Write the data to the sheet
		for i, item := range fileFilterResults {
			if i == 0 {
				for j, c := range item.Children {
					f.SetCellValue(constant.FileFilter, fmt.Sprintf("%c1", 'A'+rune(j)), c.Text)
				}
			} else {
				for j, c := range item.Children {
					f.SetCellValue(constant.FileFilter, fmt.Sprintf("%c%d", 'A'+rune(j), i+1), c.Text)
				}
			}
		}
	}

	// Create a new sheet for command
	if resultData.Spec.CommandResult != nil {
		_, err := f.NewSheet(constant.CustomCommand)
		if err != nil {
			return err
		}

		commandResults := GetCommand(resultData.Spec.CommandResult)
		// Write the data to the sheet
		for i, item := range commandResults {
			if i == 0 {
				for j, c := range item.Children {
					f.SetCellValue(constant.CustomCommand, fmt.Sprintf("%c1", 'A'+rune(j)), c.Text)
				}
			} else {
				for j, c := range item.Children {
					f.SetCellValue(constant.CustomCommand, fmt.Sprintf("%c%d", 'A'+rune(j), i+1), c.Text)
				}
			}
		}
	}

	// Create a new sheet for nodeinfo
	if resultData.Spec.NodeInfo != nil {
		_, err := f.NewSheet(constant.NodeInfo)
		if err != nil {
			return err
		}

		nodeInfoResults := GetNodeInfo(resultData.Spec.NodeInfo)
		// Write the data to the sheet
		for i, item := range nodeInfoResults {
			if i == 0 {
				for j, c := range item.Children {
					f.SetCellValue(constant.NodeInfo, fmt.Sprintf("%c1", 'A'+rune(j)), c.Text)
				}
			} else {
				for j, c := range item.Children {
					f.SetCellValue(constant.NodeInfo, fmt.Sprintf("%c%d", 'A'+rune(j), i+1), c.Text)
				}
			}
		}
	}

	// Create a new sheet for serviceconnect
	if resultData.Spec.ServiceConnectResult != nil {
		_, err := f.NewSheet(constant.ServiceConnect)
		if err != nil {
			return err
		}

		serviceConnectResults := GetServiceConnect(resultData.Spec.ServiceConnectResult)
		// Write the data to the sheet
		for i, item := range serviceConnectResults {
			if i == 0 {
				for j, c := range item.Children {
					f.SetCellValue(constant.ServiceConnect, fmt.Sprintf("%c1", 'A'+rune(j)), c.Text)
				}
			} else {
				for j, c := range item.Children {
					f.SetCellValue(constant.ServiceConnect, fmt.Sprintf("%c%d", 'A'+rune(j), i+1), c.Text)
				}
			}
		}
	}
	f.DeleteSheet("Sheet1")
	if err := f.SaveAs(fmt.Sprintf("%s.xlsx", path.Join(constant.ResultPathPrefix, resultData.Name))); err != nil {
		return err
	}

	return nil
}
