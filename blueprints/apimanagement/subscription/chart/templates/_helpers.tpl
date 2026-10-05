{{/*
Common labels applied to the Azure Service Operator Subscription resource.
*/}}
{{- define "azure-apimanagement-subscription.labels" -}}
app.kubernetes.io/managed-by: {{ .Release.Service }}
app.kubernetes.io/part-of: krateo-azure-blueprint
krateo.io/composition: {{ .Release.Name }}
{{- end -}}
