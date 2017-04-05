package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:CodigoValidacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:CodigoValidacionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:CodigoValidacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:CodigoValidacionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:CodigoValidacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:CodigoValidacionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:CodigoValidacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:CodigoValidacionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:CodigoValidacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:CodigoValidacionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:ContratoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:ContratoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:ContratoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:ContratoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:ContratoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:ContratoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:ContratoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:ContratoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:ContratoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:ContratoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:ContratoEvaluacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:ContratoEvaluacionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:ContratoEvaluacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:ContratoEvaluacionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:ContratoEvaluacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:ContratoEvaluacionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:ContratoEvaluacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:ContratoEvaluacionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:ContratoEvaluacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:ContratoEvaluacionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:ContratoProveedorController"] = append(beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:ContratoProveedorController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:ContratoProveedorController"] = append(beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:ContratoProveedorController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:ContratoProveedorController"] = append(beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:ContratoProveedorController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:ContratoProveedorController"] = append(beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:ContratoProveedorController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:ContratoProveedorController"] = append(beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:ContratoProveedorController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:DependenciaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:DependenciaController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:DependenciaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:DependenciaController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:DependenciaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:DependenciaController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:DependenciaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:DependenciaController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:DependenciaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:DependenciaController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:EvaluacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:EvaluacionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:EvaluacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:EvaluacionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:EvaluacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:EvaluacionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:EvaluacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:EvaluacionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:EvaluacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:EvaluacionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:InformacionInterventorController"] = append(beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:InformacionInterventorController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:InformacionInterventorController"] = append(beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:InformacionInterventorController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:InformacionInterventorController"] = append(beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:InformacionInterventorController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:InformacionInterventorController"] = append(beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:InformacionInterventorController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:InformacionInterventorController"] = append(beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:InformacionInterventorController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:InformacionPersonaJuridicaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:InformacionPersonaJuridicaController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:InformacionPersonaJuridicaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:InformacionPersonaJuridicaController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:InformacionPersonaJuridicaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:InformacionPersonaJuridicaController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:InformacionPersonaJuridicaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:InformacionPersonaJuridicaController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:InformacionPersonaJuridicaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:InformacionPersonaJuridicaController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:InformacionPersonaNaturalController"] = append(beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:InformacionPersonaNaturalController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:InformacionPersonaNaturalController"] = append(beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:InformacionPersonaNaturalController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:InformacionPersonaNaturalController"] = append(beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:InformacionPersonaNaturalController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:InformacionPersonaNaturalController"] = append(beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:InformacionPersonaNaturalController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:InformacionPersonaNaturalController"] = append(beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:InformacionPersonaNaturalController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:InformacionProveedorController"] = append(beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:InformacionProveedorController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:InformacionProveedorController"] = append(beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:InformacionProveedorController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:InformacionProveedorController"] = append(beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:InformacionProveedorController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:InformacionProveedorController"] = append(beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:InformacionProveedorController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:InformacionProveedorController"] = append(beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:InformacionProveedorController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:InformacionSociedadParticipanteController"] = append(beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:InformacionSociedadParticipanteController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:InformacionSociedadParticipanteController"] = append(beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:InformacionSociedadParticipanteController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:InformacionSociedadParticipanteController"] = append(beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:InformacionSociedadParticipanteController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:InformacionSociedadParticipanteController"] = append(beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:InformacionSociedadParticipanteController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:InformacionSociedadParticipanteController"] = append(beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:InformacionSociedadParticipanteController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:InformacionSociedadTemporalController"] = append(beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:InformacionSociedadTemporalController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:InformacionSociedadTemporalController"] = append(beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:InformacionSociedadTemporalController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:InformacionSociedadTemporalController"] = append(beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:InformacionSociedadTemporalController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:InformacionSociedadTemporalController"] = append(beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:InformacionSociedadTemporalController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:InformacionSociedadTemporalController"] = append(beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:InformacionSociedadTemporalController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:InformacionSupervisorController"] = append(beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:InformacionSupervisorController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:InformacionSupervisorController"] = append(beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:InformacionSupervisorController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:InformacionSupervisorController"] = append(beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:InformacionSupervisorController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:InformacionSupervisorController"] = append(beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:InformacionSupervisorController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:InformacionSupervisorController"] = append(beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:InformacionSupervisorController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:InhabilidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:InhabilidadController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:InhabilidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:InhabilidadController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:InhabilidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:InhabilidadController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:InhabilidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:InhabilidadController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:InhabilidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:InhabilidadController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:MenuController"] = append(beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:MenuController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:MenuController"] = append(beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:MenuController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:MenuController"] = append(beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:MenuController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:MenuController"] = append(beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:MenuController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:MenuController"] = append(beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:MenuController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:NomenclaturaDianController"] = append(beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:NomenclaturaDianController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:NomenclaturaDianController"] = append(beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:NomenclaturaDianController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:NomenclaturaDianController"] = append(beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:NomenclaturaDianController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:NomenclaturaDianController"] = append(beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:NomenclaturaDianController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:NomenclaturaDianController"] = append(beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:NomenclaturaDianController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:ObjetoContratarActividadCiiuController"] = append(beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:ObjetoContratarActividadCiiuController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:ObjetoContratarActividadCiiuController"] = append(beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:ObjetoContratarActividadCiiuController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:ObjetoContratarActividadCiiuController"] = append(beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:ObjetoContratarActividadCiiuController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:ObjetoContratarActividadCiiuController"] = append(beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:ObjetoContratarActividadCiiuController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:ObjetoContratarActividadCiiuController"] = append(beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:ObjetoContratarActividadCiiuController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:ObjetoContratarController"] = append(beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:ObjetoContratarController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:ObjetoContratarController"] = append(beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:ObjetoContratarController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:ObjetoContratarController"] = append(beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:ObjetoContratarController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:ObjetoContratarController"] = append(beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:ObjetoContratarController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:ObjetoContratarController"] = append(beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:ObjetoContratarController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:ObjetoContratarNucleoBasicoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:ObjetoContratarNucleoBasicoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:ObjetoContratarNucleoBasicoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:ObjetoContratarNucleoBasicoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:ObjetoContratarNucleoBasicoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:ObjetoContratarNucleoBasicoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:ObjetoContratarNucleoBasicoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:ObjetoContratarNucleoBasicoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:ObjetoContratarNucleoBasicoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:ObjetoContratarNucleoBasicoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:OrdenadorGastoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:OrdenadorGastoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:OrdenadorGastoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:OrdenadorGastoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:OrdenadorGastoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:OrdenadorGastoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:OrdenadorGastoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:OrdenadorGastoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:OrdenadorGastoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:OrdenadorGastoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:ParametroDependenciaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:ParametroDependenciaController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:ParametroDependenciaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:ParametroDependenciaController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:ParametroDependenciaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:ParametroDependenciaController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:ParametroDependenciaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:ParametroDependenciaController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:ParametroDependenciaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:ParametroDependenciaController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:ParametroEstandarController"] = append(beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:ParametroEstandarController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:ParametroEstandarController"] = append(beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:ParametroEstandarController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:ParametroEstandarController"] = append(beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:ParametroEstandarController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:ParametroEstandarController"] = append(beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:ParametroEstandarController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:ParametroEstandarController"] = append(beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:ParametroEstandarController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:ProveedorActividadCiiuController"] = append(beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:ProveedorActividadCiiuController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:ProveedorActividadCiiuController"] = append(beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:ProveedorActividadCiiuController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:ProveedorActividadCiiuController"] = append(beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:ProveedorActividadCiiuController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:ProveedorActividadCiiuController"] = append(beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:ProveedorActividadCiiuController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:ProveedorActividadCiiuController"] = append(beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:ProveedorActividadCiiuController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:ProveedorRepresentanteLegalController"] = append(beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:ProveedorRepresentanteLegalController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:ProveedorRepresentanteLegalController"] = append(beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:ProveedorRepresentanteLegalController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:ProveedorRepresentanteLegalController"] = append(beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:ProveedorRepresentanteLegalController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:ProveedorRepresentanteLegalController"] = append(beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:ProveedorRepresentanteLegalController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:ProveedorRepresentanteLegalController"] = append(beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:ProveedorRepresentanteLegalController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:ProveedorTelefonoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:ProveedorTelefonoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:ProveedorTelefonoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:ProveedorTelefonoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:ProveedorTelefonoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:ProveedorTelefonoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:ProveedorTelefonoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:ProveedorTelefonoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:ProveedorTelefonoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:ProveedorTelefonoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:SniesAreaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:SniesAreaController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:SniesAreaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:SniesAreaController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:SniesAreaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:SniesAreaController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:SniesAreaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:SniesAreaController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:SniesAreaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:SniesAreaController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:SniesNucleoBasicoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:SniesNucleoBasicoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:SniesNucleoBasicoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:SniesNucleoBasicoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:SniesNucleoBasicoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:SniesNucleoBasicoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:SniesNucleoBasicoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:SniesNucleoBasicoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:SniesNucleoBasicoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:SniesNucleoBasicoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:SolicitudCotizacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:SolicitudCotizacionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:SolicitudCotizacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:SolicitudCotizacionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:SolicitudCotizacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:SolicitudCotizacionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:SolicitudCotizacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:SolicitudCotizacionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:SolicitudCotizacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:SolicitudCotizacionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:SupervisorController"] = append(beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:SupervisorController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:SupervisorController"] = append(beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:SupervisorController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:SupervisorController"] = append(beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:SupervisorController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:SupervisorController"] = append(beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:SupervisorController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:SupervisorController"] = append(beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:SupervisorController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:TelefonoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:TelefonoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:TelefonoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:TelefonoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:TelefonoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:TelefonoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:TelefonoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:TelefonoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:TelefonoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:TelefonoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:TipoConformacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:TipoConformacionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:TipoConformacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:TipoConformacionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:TipoConformacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:TipoConformacionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:TipoConformacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:TipoConformacionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:TipoConformacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:TipoConformacionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:UnidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:UnidadController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:UnidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:UnidadController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:UnidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:UnidadController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:UnidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:UnidadController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:UnidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/agora_api_crud/controllers:UnidadController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

}
