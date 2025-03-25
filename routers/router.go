// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"github.com/udistrital/agora_api_crud/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",

		beego.NSNamespace("/codigo_validacion",
			beego.NSInclude(
				&controllers.CodigoValidacionController{},
			),
		),

		beego.NSNamespace("/informacion_persona_juridica",
			beego.NSInclude(
				&controllers.InformacionPersonaJuridicaController{},
			),
		),

		beego.NSNamespace("/contrato_evaluacion",
			beego.NSInclude(
				&controllers.ContratoEvaluacionController{},
			),
		),

		beego.NSNamespace("/informacion_persona_natural",
			beego.NSInclude(
				&controllers.InformacionPersonaNaturalController{},
			),
		),

		beego.NSNamespace("/informacion_proveedor",
			beego.NSInclude(
				&controllers.InformacionProveedorController{},
			),
		),

		beego.NSNamespace("/contrato",
			beego.NSInclude(
				&controllers.ContratoController{},
			),
		),

		beego.NSNamespace("/informacion_sociedad_temporal",
			beego.NSInclude(
				&controllers.InformacionSociedadTemporalController{},
			),
		),

		beego.NSNamespace("/menu",
			beego.NSInclude(
				&controllers.MenuController{},
			),
		),

		beego.NSNamespace("/parametro_nomenclatura_dian",
			beego.NSInclude(
				&controllers.ParametroNomenclaturaDianController{},
			),
		),

		beego.NSNamespace("/inhabilidad",
			beego.NSInclude(
				&controllers.InhabilidadController{},
			),
		),

		beego.NSNamespace("/contrato_proveedor",
			beego.NSInclude(
				&controllers.ContratoProveedorController{},
			),
		),

		beego.NSNamespace("/dependencia",
			beego.NSInclude(
				&controllers.DependenciaController{},
			),
		),

		beego.NSNamespace("/informacion_interventor",
			beego.NSInclude(
				&controllers.InformacionInterventorController{},
			),
		),

		beego.NSNamespace("/evaluacion",
			beego.NSInclude(
				&controllers.EvaluacionController{},
			),
		),

		beego.NSNamespace("/objeto_contratar_actividad_ciiu",
			beego.NSInclude(
				&controllers.ObjetoContratarActividadCiiuController{},
			),
		),

		beego.NSNamespace("/ordenador_gasto",
			beego.NSInclude(
				&controllers.OrdenadorGastoController{},
			),
		),

		beego.NSNamespace("/proveedor_actividad_ciiu",
			beego.NSInclude(
				&controllers.ProveedorActividadCiiuController{},
			),
		),

		beego.NSNamespace("/solicitud_cotizacion",
			beego.NSInclude(
				&controllers.SolicitudCotizacionController{},
			),
		),

		beego.NSNamespace("/parametro_dependencia",
			beego.NSInclude(
				&controllers.ParametroDependenciaController{},
			),
		),

		beego.NSNamespace("/proveedor_representante_legal",
			beego.NSInclude(
				&controllers.ProveedorRepresentanteLegalController{},
			),
		),

		beego.NSNamespace("/parametro_estandar",
			beego.NSInclude(
				&controllers.ParametroEstandarController{},
			),
		),

		beego.NSNamespace("/proveedor_telefono",
			beego.NSInclude(
				&controllers.ProveedorTelefonoController{},
			),
		),

		beego.NSNamespace("/snies_area",
			beego.NSInclude(
				&controllers.SniesAreaController{},
			),
		),

		beego.NSNamespace("/tipo_conformacion",
			beego.NSInclude(
				&controllers.TipoConformacionController{},
			),
		),

		beego.NSNamespace("/unidad",
			beego.NSInclude(
				&controllers.UnidadController{},
			),
		),

		beego.NSNamespace("/objeto_contratar_nucleo_basico",
			beego.NSInclude(
				&controllers.ObjetoContratarNucleoBasicoController{},
			),
		),

		beego.NSNamespace("/informacion_supervisor",
			beego.NSInclude(
				&controllers.InformacionSupervisorController{},
			),
		),

		beego.NSNamespace("/telefono",
			beego.NSInclude(
				&controllers.TelefonoController{},
			),
		),

		beego.NSNamespace("/snies_nucleo_basico",
			beego.NSInclude(
				&controllers.SniesNucleoBasicoController{},
			),
		),

		beego.NSNamespace("/objeto_contratar",
			beego.NSInclude(
				&controllers.ObjetoContratarController{},
			),
		),

		beego.NSNamespace("/informacion_sociedad_participante",
			beego.NSInclude(
				&controllers.InformacionSociedadParticipanteController{},
			),
		),

		beego.NSNamespace("/supervisor",
			beego.NSInclude(
				&controllers.SupervisorController{},
			),
		),

		beego.NSNamespace("/proveedor_contrato_persona",
		beego.NSInclude(
			&controllers.ProveedorContratoPersonaController{},
		),
	),
	)
	beego.AddNamespace(ns)
}
