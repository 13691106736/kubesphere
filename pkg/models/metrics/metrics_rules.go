/*
Copyright 2019 The KubeSphere Authors.
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
    http://www.apache.org/licenses/LICENSE-2.0
Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package metrics

import (
	"k8s.io/klog"
	"kubesphere.io/kubesphere/pkg/simple/client"
)

var (
	labelNamePod       = "pod_name"
	labelNameContainer = "container_name"
)

const (
	// TODO: expose the following metrics in prometheus format
	MetricClusterWorkspaceCount = "cluster_workspace_count"
	MetricClusterAccountCount   = "cluster_account_count"
	MetricClusterNamespaceCount = "cluster_namespace_count"
	MetricClusterDevopsCount    = "cluster_devops_project_count"

	MetricWorkspaceNamespaceCount = "workspace_namespace_count"
	MetricWorkspaceDevopsCount    = "workspace_devops_project_count"
	MetricWorkspaceMemberCount    = "workspace_member_count"
	MetricWorkspaceRoleCount      = "workspace_role_count"
)

var clusterMetrics = []string{
	"cluster_cpu_utilisation",
	"cluster_cpu_usage",
	"cluster_cpu_total",
	"cluster_memory_utilisation",
	"cluster_memory_available",
	"cluster_memory_total",
	"cluster_memory_usage_wo_cache",
	"cluster_net_utilisation",
	"cluster_net_bytes_transmitted",
	"cluster_net_bytes_received",
	"cluster_disk_read_iops",
	"cluster_disk_write_iops",
	"cluster_disk_read_throughput",
	"cluster_disk_write_throughput",
	"cluster_disk_size_usage",
	"cluster_disk_size_utilisation",
	"cluster_disk_size_capacity",
	"cluster_disk_size_available",
	"cluster_disk_inode_total",
	"cluster_disk_inode_usage",
	"cluster_disk_inode_utilisation",
	"cluster_namespace_count",
	"cluster_pod_count",
	"cluster_pod_quota",
	"cluster_pod_utilisation",
	"cluster_pod_running_count",
	"cluster_pod_succeeded_count",
	"cluster_pod_abnormal_count",
	"cluster_node_online",
	"cluster_node_offline",
	"cluster_node_total",
	"cluster_cronjob_count",
	"cluster_pvc_count",
	"cluster_daemonset_count",
	"cluster_deployment_count",
	"cluster_endpoint_count",
	"cluster_hpa_count",
	"cluster_job_count",
	"cluster_statefulset_count",
	"cluster_replicaset_count",
	"cluster_service_count",
	"cluster_secret_count",
	"cluster_pv_count",
	"cluster_ingresses_extensions_count",
	"cluster_load1",
	"cluster_load5",
	"cluster_load15",
	"cluster_pod_abnormal_ratio",
	"cluster_node_offline_ratio",
}

var nodeMetrics = []string{
	"node_cpu_utilisation",
	"node_cpu_total",
	"node_cpu_usage",
	"node_memory_utilisation",
	"node_memory_usage_wo_cache",
	"node_memory_available",
	"node_memory_total",
	"node_net_utilisation",
	"node_net_bytes_transmitted",
	"node_net_bytes_received",
	"node_disk_read_iops",
	"node_disk_write_iops",
	"node_disk_read_throughput",
	"node_disk_write_throughput",
	"node_disk_size_capacity",
	"node_disk_size_available",
	"node_disk_size_usage",
	"node_disk_size_utilisation",
	"node_disk_inode_total",
	"node_disk_inode_usage",
	"node_disk_inode_utilisation",
	"node_pod_count",
	"node_pod_quota",
	"node_pod_utilisation",
	"node_pod_running_count",
	"node_pod_succeeded_count",
	"node_pod_abnormal_count",
	"node_load1",
	"node_load5",
	"node_load15",
	"node_pod_abnormal_ratio",
}

var workspaceMetrics = []string{
	"workspace_cpu_usage",
	"workspace_memory_usage",
	"workspace_memory_usage_wo_cache",
	"workspace_net_bytes_transmitted",
	"workspace_net_bytes_received",
	"workspace_pod_count",
	"workspace_pod_running_count",
	"workspace_pod_succeeded_count",
	"workspace_pod_abnormal_count",
	"workspace_ingresses_extensions_count",
	"workspace_cronjob_count",
	"workspace_pvc_count",
	"workspace_daemonset_count",
	"workspace_deployment_count",
	"workspace_endpoint_count",
	"workspace_hpa_count",
	"workspace_job_count",
	"workspace_statefulset_count",
	"workspace_replicaset_count",
	"workspace_service_count",
	"workspace_secret_count",
	"workspace_pod_abnormal_ratio",
}

var namespaceMetrics = []string{
	"namespace_cpu_usage",
	"namespace_memory_usage",
	"namespace_memory_usage_wo_cache",
	"namespace_net_bytes_transmitted",
	"namespace_net_bytes_received",
	"namespace_pod_count",
	"namespace_pod_running_count",
	"namespace_pod_succeeded_count",
	"namespace_pod_abnormal_count",
	"namespace_pod_abnormal_ratio",
	"namespace_memory_limit_hard",
	"namespace_cpu_limit_hard",
	"namespace_pod_count_hard",
	"namespace_cronjob_count",
	"namespace_pvc_count",
	"namespace_daemonset_count",
	"namespace_deployment_count",
	"namespace_endpoint_count",
	"namespace_hpa_count",
	"namespace_job_count",
	"namespace_statefulset_count",
	"namespace_replicaset_count",
	"namespace_service_count",
	"namespace_secret_count",
	"namespace_configmap_count",
	"namespace_ingresses_extensions_count",
	"namespace_s2ibuilder_count",
}

var workloadMetrics = []string{
	// TODO: the following five metrics are deprecated.
	"workload_pod_cpu_usage",
	"workload_pod_memory_usage",
	"workload_pod_memory_usage_wo_cache",
	"workload_pod_net_bytes_transmitted",
	"workload_pod_net_bytes_received",

	"workload_cpu_usage",
	"workload_memory_usage",
	"workload_memory_usage_wo_cache",
	"workload_net_bytes_transmitted",
	"workload_net_bytes_received",

	"workload_deployment_replica",
	"workload_deployment_replica_available",
	"workload_statefulset_replica",
	"workload_statefulset_replica_available",
	"workload_daemonset_replica",
	"workload_daemonset_replica_available",
	"workload_deployment_unavailable_replicas_ratio",
	"workload_daemonset_unavailable_replicas_ratio",
	"workload_statefulset_unavailable_replicas_ratio",
}

var podMetrics = []string{
	"pod_cpu_usage",
	"pod_memory_usage",
	"pod_memory_usage_wo_cache",
	"pod_net_bytes_transmitted",
	"pod_net_bytes_received",
}

var containerMetrics = []string{
	"container_cpu_usage",
	"container_memory_usage",
	"container_memory_usage_wo_cache",
}

var pvcMetrics = []string{
	"pvc_inodes_available",
	"pvc_inodes_used",
	"pvc_inodes_total",
	"pvc_inodes_utilisation",
	"pvc_bytes_available",
	"pvc_bytes_used",
	"pvc_bytes_total",
	"pvc_bytes_utilisation",
}

var componentMetrics = []string{
	"etcd_server_list",
	"etcd_server_total",
	"etcd_server_up_total",
	"etcd_server_has_leader",
	"etcd_server_leader_changes",
	"etcd_server_proposals_failed_rate",
	"etcd_server_proposals_applied_rate",
	"etcd_server_proposals_committed_rate",
	"etcd_server_proposals_pending_count",
	"etcd_mvcc_db_size",
	"etcd_network_client_grpc_received_bytes",
	"etcd_network_client_grpc_sent_bytes",
	"etcd_grpc_call_rate",
	"etcd_grpc_call_failed_rate",
	"etcd_grpc_server_msg_received_rate",
	"etcd_grpc_server_msg_sent_rate",
	"etcd_disk_wal_fsync_duration",
	"etcd_disk_wal_fsync_duration_quantile",
	"etcd_disk_backend_commit_duration",
	"etcd_disk_backend_commit_duration_quantile",

	"apiserver_up_sum",
	"apiserver_request_rate",
	"apiserver_request_by_verb_rate",
	"apiserver_request_latencies",
	"apiserver_request_by_verb_latencies",

	"scheduler_up_sum",
	"scheduler_schedule_attempts",
	"scheduler_schedule_attempt_rate",
	"scheduler_e2e_scheduling_latency",
	"scheduler_e2e_scheduling_latency_quantile",

	"controller_manager_up_sum",

	"coredns_up_sum",
	"coredns_cache_hits",
	"coredns_cache_misses",
	"coredns_dns_request_rate",
	"coredns_dns_request_duration",
	"coredns_dns_request_duration_quantile",
	"coredns_dns_request_by_type_rate",
	"coredns_dns_request_by_rcode_rate",
	"coredns_panic_rate",
	"coredns_proxy_request_rate",
	"coredns_proxy_request_duration",
	"coredns_proxy_request_duration_quantile",

	"prometheus_up_sum",
	"prometheus_tsdb_head_samples_appended_rate",
}

var metricsPromqlMap = map[string]string{
	//cluster
	"cluster_cpu_utilisation":            ":node_cpu_utilisation:avg1m",
	"cluster_cpu_usage":                  `round(:node_cpu_utilisation:avg1m * sum(node:node_num_cpu:sum), 0.001)`,
	"cluster_cpu_total":                  "sum(node:node_num_cpu:sum)",
	"cluster_memory_utilisation":         ":node_memory_utilisation:",
	"cluster_memory_available":           "sum(node:node_memory_bytes_available:sum)",
	"cluster_memory_total":               "sum(node:node_memory_bytes_total:sum)",
	"cluster_memory_usage_wo_cache":      "sum(node:node_memory_bytes_total:sum) - sum(node:node_memory_bytes_available:sum)",
	"cluster_net_utilisation":            ":node_net_utilisation:sum_irate",
	"cluster_net_bytes_transmitted":      "sum(node:node_net_bytes_transmitted:sum_irate)",
	"cluster_net_bytes_received":         "sum(node:node_net_bytes_received:sum_irate)",
	"cluster_disk_read_iops":             "sum(node:data_volume_iops_reads:sum)",
	"cluster_disk_write_iops":            "sum(node:data_volume_iops_writes:sum)",
	"cluster_disk_read_throughput":       "sum(node:data_volume_throughput_bytes_read:sum)",
	"cluster_disk_write_throughput":      "sum(node:data_volume_throughput_bytes_written:sum)",
	"cluster_disk_size_usage":            `sum(max(node_filesystem_size_bytes{device=~"/dev/.*", device!~"/dev/loop\\d+", job="node-exporter"} - node_filesystem_avail_bytes{device=~"/dev/.*", device!~"/dev/loop\\d+", job="node-exporter"}) by (device, instance))`,
	"cluster_disk_size_utilisation":      `cluster:disk_utilization:ratio`,
	"cluster_disk_size_capacity":         `sum(max(node_filesystem_size_bytes{device=~"/dev/.*", device!~"/dev/loop\\d+", job="node-exporter"}) by (device, instance))`,
	"cluster_disk_size_available":        `sum(max(node_filesystem_avail_bytes{device=~"/dev/.*", device!~"/dev/loop\\d+", job="node-exporter"}) by (device, instance))`,
	"cluster_disk_inode_total":           `sum(node:node_inodes_total:)`,
	"cluster_disk_inode_usage":           `sum(node:node_inodes_total:) - sum(node:node_inodes_free:)`,
	"cluster_disk_inode_utilisation":     `cluster:disk_inode_utilization:ratio`,
	"cluster_namespace_count":            `count(kube_namespace_annotations)`,
	"cluster_pod_count":                  `cluster:pod:sum`,
	"cluster_pod_quota":                  `sum(max(kube_node_status_capacity_pods) by (node) unless on (node) (kube_node_status_condition{condition="Ready",status=~"unknown|false"} > 0))`,
	"cluster_pod_utilisation":            `cluster:pod_utilization:ratio`,
	"cluster_pod_running_count":          `cluster:pod_running:count`,
	"cluster_pod_succeeded_count":        `count(kube_pod_info unless on (pod) (kube_pod_status_phase{phase=~"Failed|Pending|Unknown|Running"} > 0) unless on (node) (kube_node_status_condition{condition="Ready",status=~"unknown|false"} > 0))`,
	"cluster_pod_abnormal_count":         `cluster:pod_abnormal:sum`,
	"cluster_node_online":                `sum(kube_node_status_condition{condition="Ready",status="true"})`,
	"cluster_node_offline":               `cluster:node_offline:sum`,
	"cluster_node_total":                 `sum(kube_node_status_condition{condition="Ready"})`,
	"cluster_cronjob_count":              `sum(kube_cronjob_labels)`,
	"cluster_pvc_count":                  `sum(kube_persistentvolumeclaim_info)`,
	"cluster_daemonset_count":            `sum(kube_daemonset_labels)`,
	"cluster_deployment_count":           `sum(kube_deployment_labels)`,
	"cluster_endpoint_count":             `sum(kube_endpoint_labels)`,
	"cluster_hpa_count":                  `sum(kube_hpa_labels)`,
	"cluster_job_count":                  `sum(kube_job_labels)`,
	"cluster_statefulset_count":          `sum(kube_statefulset_labels)`,
	"cluster_replicaset_count":           `count(kube_replicaset_created)`,
	"cluster_service_count":              `sum(kube_service_info)`,
	"cluster_secret_count":               `sum(kube_secret_info)`,
	"cluster_pv_count":                   `sum(kube_persistentvolume_labels)`,
	"cluster_ingresses_extensions_count": `sum(kube_ingress_labels)`,
	"cluster_load1":                      `sum(node_load1{job="node-exporter"}) / sum(node:node_num_cpu:sum)`,
	"cluster_load5":                      `sum(node_load5{job="node-exporter"}) / sum(node:node_num_cpu:sum)`,
	"cluster_load15":                     `sum(node_load15{job="node-exporter"}) / sum(node:node_num_cpu:sum)`,
	"cluster_pod_abnormal_ratio":         `cluster:pod_abnormal:ratio`,
	"cluster_node_offline_ratio":         `cluster:node_offline:ratio`,

	//node
	"node_cpu_utilisation":        "node:node_cpu_utilisation:avg1m{$1}",
	"node_cpu_total":              "node:node_num_cpu:sum{$1}",
	"node_memory_utilisation":     "node:node_memory_utilisation:{$1}",
	"node_memory_available":       "node:node_memory_bytes_available:sum{$1}",
	"node_memory_total":           "node:node_memory_bytes_total:sum{$1}",
	"node_memory_usage_wo_cache":  "node:node_memory_bytes_total:sum{$1} - node:node_memory_bytes_available:sum{$1}",
	"node_net_utilisation":        "node:node_net_utilisation:sum_irate{$1}",
	"node_net_bytes_transmitted":  "node:node_net_bytes_transmitted:sum_irate{$1}",
	"node_net_bytes_received":     "node:node_net_bytes_received:sum_irate{$1}",
	"node_disk_read_iops":         "node:data_volume_iops_reads:sum{$1}",
	"node_disk_write_iops":        "node:data_volume_iops_writes:sum{$1}",
	"node_disk_read_throughput":   "node:data_volume_throughput_bytes_read:sum{$1}",
	"node_disk_write_throughput":  "node:data_volume_throughput_bytes_written:sum{$1}",
	"node_disk_size_capacity":     `sum(max(node_filesystem_size_bytes{device=~"/dev/.*", device!~"/dev/loop\\d+", job="node-exporter"} * on (namespace, pod) group_left(node) node_namespace_pod:kube_pod_info:{$1}) by (device, node)) by (node)`,
	"node_disk_size_available":    `node:disk_space_available:{$1}`,
	"node_disk_size_usage":        `sum(max((node_filesystem_size_bytes{device=~"/dev/.*", device!~"/dev/loop\\d+", job="node-exporter"} - node_filesystem_avail_bytes{device=~"/dev/.*", device!~"/dev/loop\\d+", job="node-exporter"}) * on (namespace, pod) group_left(node) node_namespace_pod:kube_pod_info:{$1}) by (device, node)) by (node)`,
	"node_disk_size_utilisation":  `node:disk_space_utilization:ratio{$1}`,
	"node_disk_inode_total":       `node:node_inodes_total:{$1}`,
	"node_disk_inode_usage":       `node:node_inodes_total:{$1} - node:node_inodes_free:{$1}`,
	"node_disk_inode_utilisation": `node:disk_inode_utilization:ratio{$1}`,
	"node_pod_count":              `node:pod_count:sum{$1}`,
	"node_pod_quota":              `max(kube_node_status_capacity_pods{$1}) by (node) unless on (node) (kube_node_status_condition{condition="Ready",status=~"unknown|false"} > 0)`,
	"node_pod_utilisation":        `node:pod_utilization:ratio{$1}`,
	"node_pod_running_count":      `node:pod_running:count{$1}`,
	"node_pod_succeeded_count":    `node:pod_succeeded:count{$1}`,
	"node_pod_abnormal_count":     `node:pod_abnormal:count{$1}`,
	"node_cpu_usage":              `round(node:node_cpu_utilisation:avg1m{$1} * node:node_num_cpu:sum{$1}, 0.001)`,
	"node_load1":                  `node:load1:ratio{$1}`,
	"node_load5":                  `node:load5:ratio{$1}`,
	"node_load15":                 `node:load15:ratio{$1}`,
	"node_pod_abnormal_ratio":     `node:pod_abnormal:ratio{$1}`,

	// workspace
	"workspace_cpu_usage":                  `round(sum by (label_kubesphere_io_workspace) (namespace:container_cpu_usage_seconds_total:sum_rate{namespace!="", $1}), 0.001)`,
	"workspace_memory_usage":               `sum by (label_kubesphere_io_workspace) (namespace:container_memory_usage_bytes:sum{namespace!="", $1})`,
	"workspace_memory_usage_wo_cache":      `sum by (label_kubesphere_io_workspace) (namespace:container_memory_usage_bytes_wo_cache:sum{namespace!="", $1})`,
	"workspace_net_bytes_transmitted":      `sum by (label_kubesphere_io_workspace) (sum by (namespace) (irate(container_network_transmit_bytes_total{namespace!="", pod_name!="", interface!~"^(cali.+|tunl.+|dummy.+|kube.+|flannel.+|cni.+|docker.+|veth.+|lo.*)", job="kubelet"}[5m])) * on (namespace) group_left(label_kubesphere_io_workspace) kube_namespace_labels{$1})`,
	"workspace_net_bytes_received":         `sum by (label_kubesphere_io_workspace) (sum by (namespace) (irate(container_network_receive_bytes_total{namespace!="", pod_name!="", interface!~"^(cali.+|tunl.+|dummy.+|kube.+|flannel.+|cni.+|docker.+|veth.+|lo.*)", job="kubelet"}[5m])) * on (namespace) group_left(label_kubesphere_io_workspace) kube_namespace_labels{$1})`,
	"workspace_pod_count":                  `sum by (label_kubesphere_io_workspace) (kube_pod_status_phase{phase!~"Failed|Succeeded", namespace!=""} * on (namespace) group_left(label_kubesphere_io_workspace)(kube_namespace_labels{$1}))`,
	"workspace_pod_running_count":          `sum by (label_kubesphere_io_workspace) (kube_pod_status_phase{phase="Running", namespace!=""} * on (namespace) group_left(label_kubesphere_io_workspace)(kube_namespace_labels{$1}))`,
	"workspace_pod_succeeded_count":        `sum by (label_kubesphere_io_workspace) (kube_pod_status_phase{phase="Succeeded", namespace!=""} * on (namespace) group_left(label_kubesphere_io_workspace)(kube_namespace_labels{$1}))`,
	"workspace_pod_abnormal_count":         `count by (label_kubesphere_io_workspace) ((kube_pod_info{node!=""} unless on (pod, namespace) (kube_pod_status_phase{job="kube-state-metrics", phase="Succeeded"}>0) unless on (pod, namespace) ((kube_pod_status_ready{job="kube-state-metrics", condition="true"}>0) and on (pod, namespace) (kube_pod_status_phase{job="kube-state-metrics", phase="Running"}>0)) unless on (pod, namespace) (kube_pod_container_status_waiting_reason{job="kube-state-metrics", reason="ContainerCreating"}>0)) * on (namespace) group_left(label_kubesphere_io_workspace)(kube_namespace_labels{$1}))`,
	"workspace_ingresses_extensions_count": `sum by (label_kubesphere_io_workspace) (kube_ingress_labels{namespace!=""} * on (namespace) group_left(label_kubesphere_io_workspace)(kube_namespace_labels{$1}))`,
	"workspace_cronjob_count":              `sum by (label_kubesphere_io_workspace) (kube_cronjob_labels{namespace!=""} * on (namespace) group_left(label_kubesphere_io_workspace)(kube_namespace_labels{$1}))`,
	"workspace_pvc_count":                  `sum by (label_kubesphere_io_workspace) (kube_persistentvolumeclaim_info{namespace!=""} * on (namespace) group_left(label_kubesphere_io_workspace)(kube_namespace_labels{$1}))`,
	"workspace_daemonset_count":            `sum by (label_kubesphere_io_workspace) (kube_daemonset_labels{namespace!=""} * on (namespace) group_left(label_kubesphere_io_workspace)(kube_namespace_labels{$1}))`,
	"workspace_deployment_count":           `sum by (label_kubesphere_io_workspace) (kube_deployment_labels{namespace!=""} * on (namespace) group_left(label_kubesphere_io_workspace)(kube_namespace_labels{$1}))`,
	"workspace_endpoint_count":             `sum by (label_kubesphere_io_workspace) (kube_endpoint_labels{namespace!=""} * on (namespace) group_left(label_kubesphere_io_workspace)(kube_namespace_labels{$1}))`,
	"workspace_hpa_count":                  `sum by (label_kubesphere_io_workspace) (kube_hpa_labels{namespace!=""} * on (namespace) group_left(label_kubesphere_io_workspace)(kube_namespace_labels{$1}))`,
	"workspace_job_count":                  `sum by (label_kubesphere_io_workspace) (kube_job_labels{namespace!=""} * on (namespace) group_left(label_kubesphere_io_workspace)(kube_namespace_labels{$1}))`,
	"workspace_statefulset_count":          `sum by (label_kubesphere_io_workspace) (kube_statefulset_labels{namespace!=""} * on (namespace) group_left(label_kubesphere_io_workspace)(kube_namespace_labels{$1}))`,
	"workspace_replicaset_count":           `count by (label_kubesphere_io_workspace) (kube_replicaset_created{namespace!=""} * on (namespace) group_left(label_kubesphere_io_workspace)(kube_namespace_labels{$1}))`,
	"workspace_service_count":              `sum by (label_kubesphere_io_workspace) (kube_service_info{namespace!=""} * on (namespace) group_left(label_kubesphere_io_workspace)(kube_namespace_labels{$1}))`,
	"workspace_secret_count":               `sum by (label_kubesphere_io_workspace) (kube_secret_info{namespace!=""} * on (namespace) group_left(label_kubesphere_io_workspace)(kube_namespace_labels{$1}))`,
	"workspace_pod_abnormal_ratio":         `count by (label_kubesphere_io_workspace) ((kube_pod_info{node!=""} unless on (pod, namespace) (kube_pod_status_phase{job="kube-state-metrics", phase="Succeeded"}>0) unless on (pod, namespace) ((kube_pod_status_ready{job="kube-state-metrics", condition="true"}>0) and on (pod, namespace) (kube_pod_status_phase{job="kube-state-metrics", phase="Running"}>0)) unless on (pod, namespace) (kube_pod_container_status_waiting_reason{job="kube-state-metrics", reason="ContainerCreating"}>0)) * on (namespace) group_left(label_kubesphere_io_workspace) kube_namespace_labels{$1}) / sum by (label_kubesphere_io_workspace) (kube_pod_status_phase{phase!="Succeeded", namespace!=""} * on (namespace) group_left(label_kubesphere_io_workspace)(kube_namespace_labels{$1}))`,

	//namespace
	"namespace_cpu_usage":                  `round(namespace:container_cpu_usage_seconds_total:sum_rate{namespace!="", $1}, 0.001)`,
	"namespace_memory_usage":               `namespace:container_memory_usage_bytes:sum{namespace!="", $1}`,
	"namespace_memory_usage_wo_cache":      `namespace:container_memory_usage_bytes_wo_cache:sum{namespace!="", $1}`,
	"namespace_net_bytes_transmitted":      `sum by (namespace) (irate(container_network_transmit_bytes_total{namespace!="", pod_name!="", interface!~"^(cali.+|tunl.+|dummy.+|kube.+|flannel.+|cni.+|docker.+|veth.+|lo.*)", job="kubelet"}[5m]) * on (namespace) group_left(label_kubesphere_io_workspace) kube_namespace_labels{$1})`,
	"namespace_net_bytes_received":         `sum by (namespace) (irate(container_network_receive_bytes_total{namespace!="", pod_name!="", interface!~"^(cali.+|tunl.+|dummy.+|kube.+|flannel.+|cni.+|docker.+|veth.+|lo.*)", job="kubelet"}[5m]) * on (namespace) group_left(label_kubesphere_io_workspace) kube_namespace_labels{$1})`,
	"namespace_pod_count":                  `sum by (namespace) (kube_pod_status_phase{phase!~"Failed|Succeeded", namespace!=""} * on (namespace) group_left(label_kubesphere_io_workspace) kube_namespace_labels{$1})`,
	"namespace_pod_running_count":          `sum by (namespace) (kube_pod_status_phase{phase="Running", namespace!=""} * on (namespace) group_left(label_kubesphere_io_workspace) kube_namespace_labels{$1})`,
	"namespace_pod_succeeded_count":        `sum by (namespace) (kube_pod_status_phase{phase="Succeeded", namespace!=""} * on (namespace) group_left(label_kubesphere_io_workspace) kube_namespace_labels{$1})`,
	"namespace_pod_abnormal_count":         `namespace:pod_abnormal:count{namespace!="", $1}`,
	"namespace_pod_abnormal_ratio":         `namespace:pod_abnormal:ratio{namespace!="", $1}`,
	"namespace_memory_limit_hard":          `min by (namespace) (kube_resourcequota{resourcequota!="quota", type="hard", namespace!="", resource="limits.memory"} * on (namespace) group_left(label_kubesphere_io_workspace) kube_namespace_labels{$1})`,
	"namespace_cpu_limit_hard":             `min by (namespace) (kube_resourcequota{resourcequota!="quota", type="hard", namespace!="", resource="limits.cpu"} * on (namespace) group_left(label_kubesphere_io_workspace) kube_namespace_labels{$1})`,
	"namespace_pod_count_hard":             `min by (namespace) (kube_resourcequota{resourcequota!="quota", type="hard", namespace!="", resource="count/pods"} * on (namespace) group_left(label_kubesphere_io_workspace) kube_namespace_labels{$1})`,
	"namespace_cronjob_count":              `sum by (namespace) (kube_cronjob_labels{namespace!=""} * on (namespace) group_left(label_kubesphere_io_workspace) kube_namespace_labels{$1})`,
	"namespace_pvc_count":                  `sum by (namespace) (kube_persistentvolumeclaim_info{namespace!=""} * on (namespace) group_left(label_kubesphere_io_workspace) kube_namespace_labels{$1})`,
	"namespace_daemonset_count":            `sum by (namespace) (kube_daemonset_labels{namespace!=""} * on (namespace) group_left(label_kubesphere_io_workspace) kube_namespace_labels{$1})`,
	"namespace_deployment_count":           `sum by (namespace) (kube_deployment_labels{namespace!=""} * on (namespace) group_left(label_kubesphere_io_workspace) kube_namespace_labels{$1})`,
	"namespace_endpoint_count":             `sum by (namespace) (kube_endpoint_labels{namespace!=""} * on (namespace) group_left(label_kubesphere_io_workspace) kube_namespace_labels{$1})`,
	"namespace_hpa_count":                  `sum by (namespace) (kube_hpa_labels{namespace!=""} * on (namespace) group_left(label_kubesphere_io_workspace) kube_namespace_labels{$1})`,
	"namespace_job_count":                  `sum by (namespace) (kube_job_labels{namespace!=""} * on (namespace) group_left(label_kubesphere_io_workspace) kube_namespace_labels{$1})`,
	"namespace_statefulset_count":          `sum by (namespace) (kube_statefulset_labels{namespace!=""} * on (namespace) group_left(label_kubesphere_io_workspace) kube_namespace_labels{$1})`,
	"namespace_replicaset_count":           `count by (namespace) (kube_replicaset_created{namespace!=""} * on (namespace) group_left(label_kubesphere_io_workspace) kube_namespace_labels{$1})`,
	"namespace_service_count":              `sum by (namespace) (kube_service_info{namespace!=""} * on (namespace) group_left(label_kubesphere_io_workspace) kube_namespace_labels{$1})`,
	"namespace_secret_count":               `sum by (namespace) (kube_secret_info{namespace!=""} * on (namespace) group_left(label_kubesphere_io_workspace) kube_namespace_labels{$1})`,
	"namespace_configmap_count":            `sum by (namespace) (kube_configmap_info{namespace!=""} * on (namespace) group_left(label_kubesphere_io_workspace) kube_namespace_labels{$1})`,
	"namespace_ingresses_extensions_count": `sum by (namespace) (kube_ingress_labels{namespace!=""} * on (namespace) group_left(label_kubesphere_io_workspace) kube_namespace_labels{$1})`,
	"namespace_s2ibuilder_count":           `sum by (namespace) (s2i_s2ibuilder_created{namespace!=""} * on (namespace) group_left(label_kubesphere_io_workspace) kube_namespace_labels{$1})`,

	// workload
	// TODO: the following five metrics are deprecated.
	"workload_pod_cpu_usage":             `round(namespace:workload_cpu_usage:sum{$1}, 0.001)`,
	"workload_pod_memory_usage":          `namespace:workload_memory_usage:sum{$1}`,
	"workload_pod_memory_usage_wo_cache": `namespace:workload_memory_usage_wo_cache:sum{$1}`,
	"workload_pod_net_bytes_transmitted": `namespace:workload_net_bytes_transmitted:sum_irate{$1}`,
	"workload_pod_net_bytes_received":    `namespace:workload_net_bytes_received:sum_irate{$1}`,

	"workload_cpu_usage":             `round(namespace:workload_cpu_usage:sum{$1}, 0.001)`,
	"workload_memory_usage":          `namespace:workload_memory_usage:sum{$1}`,
	"workload_memory_usage_wo_cache": `namespace:workload_memory_usage_wo_cache:sum{$1}`,
	"workload_net_bytes_transmitted": `namespace:workload_net_bytes_transmitted:sum_irate{$1}`,
	"workload_net_bytes_received":    `namespace:workload_net_bytes_received:sum_irate{$1}`,

	"workload_deployment_replica":                     `label_join(sum (label_join(label_replace(kube_deployment_spec_replicas{$2}, "owner_kind", "Deployment", "", ""), "workload", "", "deployment")) by (namespace, owner_kind, workload), "workload", ":", "owner_kind", "workload")`,
	"workload_deployment_replica_available":           `label_join(sum (label_join(label_replace(kube_deployment_status_replicas_available{$2}, "owner_kind", "Deployment", "", ""), "workload", "", "deployment")) by (namespace, owner_kind, workload), "workload", ":", "owner_kind", "workload")`,
	"workload_statefulset_replica":                    `label_join(sum (label_join(label_replace(kube_statefulset_replicas{$2}, "owner_kind", "StatefulSet", "", ""), "workload", "", "statefulset")) by (namespace, owner_kind, workload), "workload", ":", "owner_kind", "workload")`,
	"workload_statefulset_replica_available":          `label_join(sum (label_join(label_replace(kube_statefulset_status_replicas_current{$2}, "owner_kind", "StatefulSet", "", ""), "workload", "", "statefulset")) by (namespace, owner_kind, workload), "workload", ":", "owner_kind", "workload")`,
	"workload_daemonset_replica":                      `label_join(sum (label_join(label_replace(kube_daemonset_status_desired_number_scheduled{$2}, "owner_kind", "DaemonSet", "", ""), "workload", "", "daemonset")) by (namespace, owner_kind, workload), "workload", ":", "owner_kind", "workload")`,
	"workload_daemonset_replica_available":            `label_join(sum (label_join(label_replace(kube_daemonset_status_number_available{$2}, "owner_kind", "DaemonSet", "", ""), "workload", "", "daemonset")) by (namespace, owner_kind, workload), "workload", ":", "owner_kind", "workload")`,
	"workload_deployment_unavailable_replicas_ratio":  `namespace:deployment_unavailable_replicas:ratio{$1}`,
	"workload_daemonset_unavailable_replicas_ratio":   `namespace:daemonset_unavailable_replicas:ratio{$1}`,
	"workload_statefulset_unavailable_replicas_ratio": `namespace:statefulset_unavailable_replicas:ratio{$1}`,

	// pod
	"pod_cpu_usage":             `round(label_join(sum by (namespace, pod_name) (irate(container_cpu_usage_seconds_total{job="kubelet", pod_name!="", image!=""}[5m])), "pod", "", "pod_name") * on (namespace, pod) group_left(owner_kind, owner_name) kube_pod_owner{$1} * on (namespace, pod) group_left(node) kube_pod_info{$2}, 0.001)`,
	"pod_memory_usage":          `label_join(sum by (namespace, pod_name) (container_memory_usage_bytes{job="kubelet", pod_name!="", image!=""}), "pod", "", "pod_name") * on (namespace, pod) group_left(owner_kind, owner_name) kube_pod_owner{$1} * on (namespace, pod) group_left(node) kube_pod_info{$2}`,
	"pod_memory_usage_wo_cache": `label_join(sum by (namespace, pod_name) (container_memory_working_set_bytes{job="kubelet", pod_name!="", image!=""}), "pod", "", "pod_name") * on (namespace, pod) group_left(owner_kind, owner_name) kube_pod_owner{$1} * on (namespace, pod) group_left(node) kube_pod_info{$2}`,
	"pod_net_bytes_transmitted": `label_join(sum by (namespace, pod_name) (irate(container_network_transmit_bytes_total{pod_name!="", interface!~"^(cali.+|tunl.+|dummy.+|kube.+|flannel.+|cni.+|docker.+|veth.+|lo.*)", job="kubelet"}[5m])), "pod", "", "pod_name") * on (namespace, pod) group_left(owner_kind, owner_name) kube_pod_owner{$1} * on (namespace, pod) group_left(node) kube_pod_info{$2}`,
	"pod_net_bytes_received":    `label_join(sum by (namespace, pod_name) (irate(container_network_receive_bytes_total{pod_name!="", interface!~"^(cali.+|tunl.+|dummy.+|kube.+|flannel.+|cni.+|docker.+|veth.+|lo.*)", job="kubelet"}[5m])), "pod", "", "pod_name") * on (namespace, pod) group_left(owner_kind, owner_name) kube_pod_owner{$1} * on (namespace, pod) group_left(node) kube_pod_info{$2}`,

	// container
	"container_cpu_usage":             `round(sum by (namespace, pod_name, container_name) (irate(container_cpu_usage_seconds_total{job="kubelet", container_name!="POD", container_name!="", image!="", $1}[5m])), 0.001)`,
	"container_memory_usage":          `sum by (namespace, pod_name, container_name) (container_memory_usage_bytes{job="kubelet", container_name!="POD", container_name!="", image!="", $1})`,
	"container_memory_usage_wo_cache": `sum by (namespace, pod_name, container_name) (container_memory_working_set_bytes{job="kubelet", container_name!="POD", container_name!="", image!="", $1})`,

	// pvc
	"pvc_inodes_available":   `max by (namespace, persistentvolumeclaim) (kubelet_volume_stats_inodes_free) * on (namespace, persistentvolumeclaim) group_left (storageclass) kube_persistentvolumeclaim_info{$1}`,
	"pvc_inodes_used":        `max by (namespace, persistentvolumeclaim) (kubelet_volume_stats_inodes_used) * on (namespace, persistentvolumeclaim) group_left (storageclass) kube_persistentvolumeclaim_info{$1}`,
	"pvc_inodes_total":       `max by (namespace, persistentvolumeclaim) (kubelet_volume_stats_inodes) * on (namespace, persistentvolumeclaim) group_left (storageclass) kube_persistentvolumeclaim_info{$1}`,
	"pvc_inodes_utilisation": `max by (namespace, persistentvolumeclaim) (kubelet_volume_stats_inodes_used / kubelet_volume_stats_inodes) * on (namespace, persistentvolumeclaim) group_left (storageclass) kube_persistentvolumeclaim_info{$1}`,
	"pvc_bytes_available":    `max by (namespace, persistentvolumeclaim) (kubelet_volume_stats_available_bytes) * on (namespace, persistentvolumeclaim) group_left (storageclass) kube_persistentvolumeclaim_info{$1}`,
	"pvc_bytes_used":         `max by (namespace, persistentvolumeclaim) (kubelet_volume_stats_used_bytes) * on (namespace, persistentvolumeclaim) group_left (storageclass) kube_persistentvolumeclaim_info{$1}`,
	"pvc_bytes_total":        `max by (namespace, persistentvolumeclaim) (kubelet_volume_stats_capacity_bytes) * on (namespace, persistentvolumeclaim) group_left (storageclass) kube_persistentvolumeclaim_info{$1}`,
	"pvc_bytes_utilisation":  `max by (namespace, persistentvolumeclaim) (kubelet_volume_stats_used_bytes / kubelet_volume_stats_capacity_bytes) * on (namespace, persistentvolumeclaim) group_left (storageclass) kube_persistentvolumeclaim_info{$1}`,

	// component
	"etcd_server_list":                           `label_replace(up{job="etcd"}, "node_ip", "$1", "instance", "(.*):.*")`,
	"etcd_server_total":                          `count(up{job="etcd"})`,
	"etcd_server_up_total":                       `etcd:up:sum`,
	"etcd_server_has_leader":                     `label_replace(etcd_server_has_leader, "node_ip", "$1", "instance", "(.*):.*")`,
	"etcd_server_leader_changes":                 `label_replace(etcd:etcd_server_leader_changes_seen:sum_changes, "node_ip", "$1", "node", "(.*)")`,
	"etcd_server_proposals_failed_rate":          `avg(etcd:etcd_server_proposals_failed:sum_irate)`,
	"etcd_server_proposals_applied_rate":         `avg(etcd:etcd_server_proposals_applied:sum_irate)`,
	"etcd_server_proposals_committed_rate":       `avg(etcd:etcd_server_proposals_committed:sum_irate)`,
	"etcd_server_proposals_pending_count":        `avg(etcd:etcd_server_proposals_pending:sum)`,
	"etcd_mvcc_db_size":                          `avg(etcd:etcd_debugging_mvcc_db_total_size:sum)`,
	"etcd_network_client_grpc_received_bytes":    `sum(etcd:etcd_network_client_grpc_received_bytes:sum_irate)`,
	"etcd_network_client_grpc_sent_bytes":        `sum(etcd:etcd_network_client_grpc_sent_bytes:sum_irate)`,
	"etcd_grpc_call_rate":                        `sum(etcd:grpc_server_started:sum_irate)`,
	"etcd_grpc_call_failed_rate":                 `sum(etcd:grpc_server_handled:sum_irate)`,
	"etcd_grpc_server_msg_received_rate":         `sum(etcd:grpc_server_msg_received:sum_irate)`,
	"etcd_grpc_server_msg_sent_rate":             `sum(etcd:grpc_server_msg_sent:sum_irate)`,
	"etcd_disk_wal_fsync_duration":               `avg(etcd:etcd_disk_wal_fsync_duration:avg)`,
	"etcd_disk_wal_fsync_duration_quantile":      `avg(etcd:etcd_disk_wal_fsync_duration:histogram_quantile) by (quantile)`,
	"etcd_disk_backend_commit_duration":          `avg(etcd:etcd_disk_backend_commit_duration:avg)`,
	"etcd_disk_backend_commit_duration_quantile": `avg(etcd:etcd_disk_backend_commit_duration:histogram_quantile) by (quantile)`,

	"apiserver_up_sum":                    `apiserver:up:sum`,
	"apiserver_request_rate":              `apiserver:apiserver_request_count:sum_irate`,
	"apiserver_request_by_verb_rate":      `apiserver:apiserver_request_count:sum_verb_irate`,
	"apiserver_request_latencies":         `apiserver:apiserver_request_latencies:avg`,
	"apiserver_request_by_verb_latencies": `apiserver:apiserver_request_latencies:avg_by_verb`,

	"scheduler_up_sum":                          `scheduler:up:sum`,
	"scheduler_schedule_attempts":               `scheduler:scheduler_schedule_attempts:sum`,
	"scheduler_schedule_attempt_rate":           `scheduler:scheduler_schedule_attempts:sum_rate`,
	"scheduler_e2e_scheduling_latency":          `scheduler:scheduler_e2e_scheduling_latency:avg`,
	"scheduler_e2e_scheduling_latency_quantile": `scheduler:scheduler_e2e_scheduling_latency:histogram_quantile`,

	"controller_manager_up_sum": `controller_manager:up:sum`,

	"coredns_up_sum":                          `coredns:up:sum`,
	"coredns_cache_hits":                      `coredns:coredns_cache_hits_total:sum_irate`,
	"coredns_cache_misses":                    `coredns:coredns_cache_misses:sum_irate`,
	"coredns_dns_request_rate":                `coredns:coredns_dns_request_count:sum_irate`,
	"coredns_dns_request_duration":            `coredns:coredns_dns_request_duration:avg`,
	"coredns_dns_request_duration_quantile":   `coredns:coredns_dns_request_duration:histogram_quantile`,
	"coredns_dns_request_by_type_rate":        `coredns:coredns_dns_request_type_count:sum_irate`,
	"coredns_dns_request_by_rcode_rate":       `coredns:coredns_dns_response_rcode_count:sum_irate`,
	"coredns_panic_rate":                      `coredns:coredns_panic_count:sum_irate`,
	"coredns_proxy_request_rate":              `coredns:coredns_proxy_request_count:sum_irate`,
	"coredns_proxy_request_duration":          `coredns:coredns_proxy_request_duration:avg`,
	"coredns_proxy_request_duration_quantile": `coredns:coredns_proxy_request_duration:histogram_quantile`,

	"prometheus_up_sum":                          `prometheus:up:sum`,
	"prometheus_tsdb_head_samples_appended_rate": `prometheus:prometheus_tsdb_head_samples_appended:sum_rate`,
}

// As of Kubernetes v1.16, any Prometheus queries that match `pod_name` and
// `container_name` labels must be updated to use `pod` and `container` instead.
func CompatibleMetrics() {
	version, err := client.ClientSets().K8s().Discovery().ServerVersion()
	if err != nil {
		klog.Errorf("fail to fetch k8s version: %v", err)
		return
	}

	if version.Minor >= "16" {
		labelNamePod = "pod"
		labelNameContainer = "container"

		m := metricsPromqlMap
		m["workspace_net_bytes_transmitted"] = `sum by (label_kubesphere_io_workspace) (sum by (namespace) (irate(container_network_transmit_bytes_total{namespace!="", pod!="", interface!~"^(cali.+|tunl.+|dummy.+|kube.+|flannel.+|cni.+|docker.+|veth.+|lo.*)", job="kubelet"}[5m])) * on (namespace) group_left(label_kubesphere_io_workspace) kube_namespace_labels{$1})`
		m["workspace_net_bytes_received"] = `sum by (label_kubesphere_io_workspace) (sum by (namespace) (irate(container_network_receive_bytes_total{namespace!="", pod!="", interface!~"^(cali.+|tunl.+|dummy.+|kube.+|flannel.+|cni.+|docker.+|veth.+|lo.*)", job="kubelet"}[5m])) * on (namespace) group_left(label_kubesphere_io_workspace) kube_namespace_labels{$1})`
		m["namespace_net_bytes_transmitted"] = `sum by (namespace) (irate(container_network_transmit_bytes_total{namespace!="", pod!="", interface!~"^(cali.+|tunl.+|dummy.+|kube.+|flannel.+|cni.+|docker.+|veth.+|lo.*)", job="kubelet"}[5m]) * on (namespace) group_left(label_kubesphere_io_workspace) kube_namespace_labels{$1})`
		m["namespace_net_bytes_received"] = `sum by (namespace) (irate(container_network_receive_bytes_total{namespace!="", pod!="", interface!~"^(cali.+|tunl.+|dummy.+|kube.+|flannel.+|cni.+|docker.+|veth.+|lo.*)", job="kubelet"}[5m]) * on (namespace) group_left(label_kubesphere_io_workspace) kube_namespace_labels{$1})`
		m["pod_cpu_usage"] = `round(sum by (namespace, pod) (irate(container_cpu_usage_seconds_total{job="kubelet", pod!="", image!=""}[5m])) * on (namespace, pod) group_left(owner_kind, owner_name) kube_pod_owner{$1} * on (namespace, pod) group_left(node) kube_pod_info{$2}, 0.001)`
		m["pod_memory_usage"] = `sum by (namespace, pod) (container_memory_usage_bytes{job="kubelet", pod!="", image!=""}) * on (namespace, pod) group_left(owner_kind, owner_name) kube_pod_owner{$1} * on (namespace, pod) group_left(node) kube_pod_info{$2}`
		m["pod_memory_usage_wo_cache"] = `sum by (namespace, pod) (container_memory_working_set_bytes{job="kubelet", pod!="", image!=""}) * on (namespace, pod) group_left(owner_kind, owner_name) kube_pod_owner{$1} * on (namespace, pod) group_left(node) kube_pod_info{$2}`
		m["pod_net_bytes_transmitted"] = `sum by (namespace, pod) (irate(container_network_transmit_bytes_total{pod!="", interface!~"^(cali.+|tunl.+|dummy.+|kube.+|flannel.+|cni.+|docker.+|veth.+|lo.*)", job="kubelet"}[5m])) * on (namespace, pod) group_left(owner_kind, owner_name) kube_pod_owner{$1} * on (namespace, pod) group_left(node) kube_pod_info{$2}`
		m["pod_net_bytes_received"] = `sum by (namespace, pod) (irate(container_network_receive_bytes_total{pod!="", interface!~"^(cali.+|tunl.+|dummy.+|kube.+|flannel.+|cni.+|docker.+|veth.+|lo.*)", job="kubelet"}[5m])) * on (namespace, pod) group_left(owner_kind, owner_name) kube_pod_owner{$1} * on (namespace, pod) group_left(node) kube_pod_info{$2}`
		m["container_cpu_usage"] = `round(sum by (namespace, pod, container) (irate(container_cpu_usage_seconds_total{job="kubelet", container!="POD", container!="", image!="", $1}[5m])), 0.001)`
		m["container_memory_usage"] = `sum by (namespace, pod, container) (container_memory_usage_bytes{job="kubelet", container!="POD", container!="", image!="", $1})`
		m["container_memory_usage_wo_cache"] = `sum by (namespace, pod, container) (container_memory_working_set_bytes{job="kubelet", container!="POD", container!="", image!="", $1})`
	}
}
