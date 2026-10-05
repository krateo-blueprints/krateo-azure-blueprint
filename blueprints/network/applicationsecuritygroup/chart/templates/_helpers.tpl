{{/*
Common labels applied to the Azure Service Operator ApplicationSecurityGroup resource.
*/}}
{{- define "azure-network-applicationsecuritygroup.labels" -}}
app.kubernetes.io/managed-by: {{ .Release.Service }}
app.kubernetes.io/part-of: krateo-azure-blueprint
krateo.io/composition: {{ .Release.Name }}
{{- end -}}
