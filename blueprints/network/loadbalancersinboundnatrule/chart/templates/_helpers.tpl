{{/*
Common labels applied to the Azure Service Operator LoadBalancersInboundNatRule resource.
*/}}
{{- define "azure-network-loadbalancersinboundnatrule.labels" -}}
app.kubernetes.io/managed-by: {{ .Release.Service }}
app.kubernetes.io/part-of: krateo-azure-blueprint
krateo.io/composition: {{ .Release.Name }}
{{- end -}}
