/* tslint:disable */
/* eslint-disable */
/**
 * SUT SA Example API Playlist Vidoe
 * This is a sample server for SUT SE 2563
 *
 * The version of the OpenAPI document: 1.0
 * Contact: support@swagger.io
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */


import * as runtime from '../runtime';
import {
    ControllersDiagnose,
    ControllersDiagnoseFromJSON,
    ControllersDiagnoseToJSON,
    EntDiagnose,
    EntDiagnoseFromJSON,
    EntDiagnoseToJSON,
    EntDoctor,
    EntDoctorFromJSON,
    EntDoctorToJSON,
    EntIllness,
    EntIllnessFromJSON,
    EntIllnessToJSON,
    EntPatient,
    EntPatientFromJSON,
    EntPatientToJSON,
    EntSymptoms,
    EntSymptomsFromJSON,
    EntSymptomsToJSON,
} from '../models';

export interface CreateDiagnoseRequest {
    diagnose: ControllersDiagnose;
}

export interface CreateDoctorRequest {
    doctor: EntDoctor;
}

export interface CreateIllnessRequest {
    illness: EntIllness;
}

export interface CreatePatientRequest {
    patient: EntPatient;
}

export interface CreateSymptomsRequest {
    symptoms: EntSymptoms;
}

export interface DeleteDiagnoseRequest {
    id: number;
}

export interface DeleteDoctorRequest {
    id: number;
}

export interface DeleteIllnessRequest {
    id: number;
}

export interface DeletePatientRequest {
    id: number;
}

export interface DeleteSymptomsRequest {
    id: number;
}

export interface GetDoctorRequest {
    id: number;
}

export interface GetDrugAllergyRequest {
    id: number;
}

export interface GetIllnessRequest {
    id: number;
}

export interface GetPatientRequest {
    id: number;
}

export interface GetSymptomsRequest {
    id: number;
}

export interface ListDiagnoseRequest {
    limit?: number;
    offset?: number;
}

export interface ListDoctorRequest {
    limit?: number;
    offset?: number;
}

export interface ListIllnessRequest {
    limit?: number;
    offset?: number;
}

export interface ListPatientRequest {
    limit?: number;
    offset?: number;
}

export interface ListSymptomsRequest {
    limit?: number;
    offset?: number;
}

export interface UpdateDiagnoseRequest {
    id: number;
    drugAllergy: EntDiagnose;
}

export interface UpdateDoctorRequest {
    id: number;
    doctor: EntDoctor;
}

export interface UpdateIllnessRequest {
    id: number;
    illness: EntIllness;
}

export interface UpdatePatientRequest {
    id: number;
    patient: EntPatient;
}

export interface UpdateSymptomsRequest {
    id: number;
    symptoms: EntSymptoms;
}

/**
 * 
 */
export class DefaultApi extends runtime.BaseAPI {

    /**
     * Create diagnose
     * Create diagnose
     */
    async createDiagnoseRaw(requestParameters: CreateDiagnoseRequest): Promise<runtime.ApiResponse<EntDiagnose>> {
        if (requestParameters.diagnose === null || requestParameters.diagnose === undefined) {
            throw new runtime.RequiredError('diagnose','Required parameter requestParameters.diagnose was null or undefined when calling createDiagnose.');
        }

        const queryParameters: runtime.HTTPQuery = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        const response = await this.request({
            path: `/diagnoses`,
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
            body: ControllersDiagnoseToJSON(requestParameters.diagnose),
        });

        return new runtime.JSONApiResponse(response, (jsonValue) => EntDiagnoseFromJSON(jsonValue));
    }

    /**
     * Create diagnose
     * Create diagnose
     */
    async createDiagnose(requestParameters: CreateDiagnoseRequest): Promise<EntDiagnose> {
        const response = await this.createDiagnoseRaw(requestParameters);
        return await response.value();
    }

    /**
     * Create doctor
     * Create doctor
     */
    async createDoctorRaw(requestParameters: CreateDoctorRequest): Promise<runtime.ApiResponse<EntDoctor>> {
        if (requestParameters.doctor === null || requestParameters.doctor === undefined) {
            throw new runtime.RequiredError('doctor','Required parameter requestParameters.doctor was null or undefined when calling createDoctor.');
        }

        const queryParameters: runtime.HTTPQuery = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        const response = await this.request({
            path: `/doctors`,
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
            body: EntDoctorToJSON(requestParameters.doctor),
        });

        return new runtime.JSONApiResponse(response, (jsonValue) => EntDoctorFromJSON(jsonValue));
    }

    /**
     * Create doctor
     * Create doctor
     */
    async createDoctor(requestParameters: CreateDoctorRequest): Promise<EntDoctor> {
        const response = await this.createDoctorRaw(requestParameters);
        return await response.value();
    }

    /**
     * Create illness
     * Create illness
     */
    async createIllnessRaw(requestParameters: CreateIllnessRequest): Promise<runtime.ApiResponse<EntIllness>> {
        if (requestParameters.illness === null || requestParameters.illness === undefined) {
            throw new runtime.RequiredError('illness','Required parameter requestParameters.illness was null or undefined when calling createIllness.');
        }

        const queryParameters: runtime.HTTPQuery = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        const response = await this.request({
            path: `/illnesss`,
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
            body: EntIllnessToJSON(requestParameters.illness),
        });

        return new runtime.JSONApiResponse(response, (jsonValue) => EntIllnessFromJSON(jsonValue));
    }

    /**
     * Create illness
     * Create illness
     */
    async createIllness(requestParameters: CreateIllnessRequest): Promise<EntIllness> {
        const response = await this.createIllnessRaw(requestParameters);
        return await response.value();
    }

    /**
     * Create patient
     * Create patient
     */
    async createPatientRaw(requestParameters: CreatePatientRequest): Promise<runtime.ApiResponse<EntPatient>> {
        if (requestParameters.patient === null || requestParameters.patient === undefined) {
            throw new runtime.RequiredError('patient','Required parameter requestParameters.patient was null or undefined when calling createPatient.');
        }

        const queryParameters: runtime.HTTPQuery = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        const response = await this.request({
            path: `/patients`,
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
            body: EntPatientToJSON(requestParameters.patient),
        });

        return new runtime.JSONApiResponse(response, (jsonValue) => EntPatientFromJSON(jsonValue));
    }

    /**
     * Create patient
     * Create patient
     */
    async createPatient(requestParameters: CreatePatientRequest): Promise<EntPatient> {
        const response = await this.createPatientRaw(requestParameters);
        return await response.value();
    }

    /**
     * Create symptoms
     * Create symptoms
     */
    async createSymptomsRaw(requestParameters: CreateSymptomsRequest): Promise<runtime.ApiResponse<EntSymptoms>> {
        if (requestParameters.symptoms === null || requestParameters.symptoms === undefined) {
            throw new runtime.RequiredError('symptoms','Required parameter requestParameters.symptoms was null or undefined when calling createSymptoms.');
        }

        const queryParameters: runtime.HTTPQuery = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        const response = await this.request({
            path: `/symptomss`,
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
            body: EntSymptomsToJSON(requestParameters.symptoms),
        });

        return new runtime.JSONApiResponse(response, (jsonValue) => EntSymptomsFromJSON(jsonValue));
    }

    /**
     * Create symptoms
     * Create symptoms
     */
    async createSymptoms(requestParameters: CreateSymptomsRequest): Promise<EntSymptoms> {
        const response = await this.createSymptomsRaw(requestParameters);
        return await response.value();
    }

    /**
     * get diagnose by ID
     * Delete a diagnose entity by ID
     */
    async deleteDiagnoseRaw(requestParameters: DeleteDiagnoseRequest): Promise<runtime.ApiResponse<object>> {
        if (requestParameters.id === null || requestParameters.id === undefined) {
            throw new runtime.RequiredError('id','Required parameter requestParameters.id was null or undefined when calling deleteDiagnose.');
        }

        const queryParameters: runtime.HTTPQuery = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/diagnoses/{id}`.replace(`{${"id"}}`, encodeURIComponent(String(requestParameters.id))),
            method: 'DELETE',
            headers: headerParameters,
            query: queryParameters,
        });

        return new runtime.JSONApiResponse<any>(response);
    }

    /**
     * get diagnose by ID
     * Delete a diagnose entity by ID
     */
    async deleteDiagnose(requestParameters: DeleteDiagnoseRequest): Promise<object> {
        const response = await this.deleteDiagnoseRaw(requestParameters);
        return await response.value();
    }

    /**
     * get doctor by ID
     * Delete a doctor entity by ID
     */
    async deleteDoctorRaw(requestParameters: DeleteDoctorRequest): Promise<runtime.ApiResponse<object>> {
        if (requestParameters.id === null || requestParameters.id === undefined) {
            throw new runtime.RequiredError('id','Required parameter requestParameters.id was null or undefined when calling deleteDoctor.');
        }

        const queryParameters: runtime.HTTPQuery = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/doctors/{id}`.replace(`{${"id"}}`, encodeURIComponent(String(requestParameters.id))),
            method: 'DELETE',
            headers: headerParameters,
            query: queryParameters,
        });

        return new runtime.JSONApiResponse<any>(response);
    }

    /**
     * get doctor by ID
     * Delete a doctor entity by ID
     */
    async deleteDoctor(requestParameters: DeleteDoctorRequest): Promise<object> {
        const response = await this.deleteDoctorRaw(requestParameters);
        return await response.value();
    }

    /**
     * get illness by ID
     * Delete a illness entity by ID
     */
    async deleteIllnessRaw(requestParameters: DeleteIllnessRequest): Promise<runtime.ApiResponse<object>> {
        if (requestParameters.id === null || requestParameters.id === undefined) {
            throw new runtime.RequiredError('id','Required parameter requestParameters.id was null or undefined when calling deleteIllness.');
        }

        const queryParameters: runtime.HTTPQuery = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/illnesss/{id}`.replace(`{${"id"}}`, encodeURIComponent(String(requestParameters.id))),
            method: 'DELETE',
            headers: headerParameters,
            query: queryParameters,
        });

        return new runtime.JSONApiResponse<any>(response);
    }

    /**
     * get illness by ID
     * Delete a illness entity by ID
     */
    async deleteIllness(requestParameters: DeleteIllnessRequest): Promise<object> {
        const response = await this.deleteIllnessRaw(requestParameters);
        return await response.value();
    }

    /**
     * get patient by ID
     * Delete a patient entity by ID
     */
    async deletePatientRaw(requestParameters: DeletePatientRequest): Promise<runtime.ApiResponse<object>> {
        if (requestParameters.id === null || requestParameters.id === undefined) {
            throw new runtime.RequiredError('id','Required parameter requestParameters.id was null or undefined when calling deletePatient.');
        }

        const queryParameters: runtime.HTTPQuery = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/patients/{id}`.replace(`{${"id"}}`, encodeURIComponent(String(requestParameters.id))),
            method: 'DELETE',
            headers: headerParameters,
            query: queryParameters,
        });

        return new runtime.JSONApiResponse<any>(response);
    }

    /**
     * get patient by ID
     * Delete a patient entity by ID
     */
    async deletePatient(requestParameters: DeletePatientRequest): Promise<object> {
        const response = await this.deletePatientRaw(requestParameters);
        return await response.value();
    }

    /**
     * get symptoms by ID
     * Delete a symptoms entity by ID
     */
    async deleteSymptomsRaw(requestParameters: DeleteSymptomsRequest): Promise<runtime.ApiResponse<object>> {
        if (requestParameters.id === null || requestParameters.id === undefined) {
            throw new runtime.RequiredError('id','Required parameter requestParameters.id was null or undefined when calling deleteSymptoms.');
        }

        const queryParameters: runtime.HTTPQuery = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/symptomss/{id}`.replace(`{${"id"}}`, encodeURIComponent(String(requestParameters.id))),
            method: 'DELETE',
            headers: headerParameters,
            query: queryParameters,
        });

        return new runtime.JSONApiResponse<any>(response);
    }

    /**
     * get symptoms by ID
     * Delete a symptoms entity by ID
     */
    async deleteSymptoms(requestParameters: DeleteSymptomsRequest): Promise<object> {
        const response = await this.deleteSymptomsRaw(requestParameters);
        return await response.value();
    }

    /**
     * get doctor by ID
     * Get a doctor entity by ID
     */
    async getDoctorRaw(requestParameters: GetDoctorRequest): Promise<runtime.ApiResponse<EntDoctor>> {
        if (requestParameters.id === null || requestParameters.id === undefined) {
            throw new runtime.RequiredError('id','Required parameter requestParameters.id was null or undefined when calling getDoctor.');
        }

        const queryParameters: runtime.HTTPQuery = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/doctors/{id}`.replace(`{${"id"}}`, encodeURIComponent(String(requestParameters.id))),
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        });

        return new runtime.JSONApiResponse(response, (jsonValue) => EntDoctorFromJSON(jsonValue));
    }

    /**
     * get doctor by ID
     * Get a doctor entity by ID
     */
    async getDoctor(requestParameters: GetDoctorRequest): Promise<EntDoctor> {
        const response = await this.getDoctorRaw(requestParameters);
        return await response.value();
    }

    /**
     * get diagnose by ID
     * Get a diagnose entity by ID
     */
    async getDrugAllergyRaw(requestParameters: GetDrugAllergyRequest): Promise<runtime.ApiResponse<EntDiagnose>> {
        if (requestParameters.id === null || requestParameters.id === undefined) {
            throw new runtime.RequiredError('id','Required parameter requestParameters.id was null or undefined when calling getDrugAllergy.');
        }

        const queryParameters: runtime.HTTPQuery = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/diagnoses/{id}`.replace(`{${"id"}}`, encodeURIComponent(String(requestParameters.id))),
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        });

        return new runtime.JSONApiResponse(response, (jsonValue) => EntDiagnoseFromJSON(jsonValue));
    }

    /**
     * get diagnose by ID
     * Get a diagnose entity by ID
     */
    async getDrugAllergy(requestParameters: GetDrugAllergyRequest): Promise<EntDiagnose> {
        const response = await this.getDrugAllergyRaw(requestParameters);
        return await response.value();
    }

    /**
     * get illness by ID
     * Get a illness entity by ID
     */
    async getIllnessRaw(requestParameters: GetIllnessRequest): Promise<runtime.ApiResponse<EntIllness>> {
        if (requestParameters.id === null || requestParameters.id === undefined) {
            throw new runtime.RequiredError('id','Required parameter requestParameters.id was null or undefined when calling getIllness.');
        }

        const queryParameters: runtime.HTTPQuery = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/illnesss/{id}`.replace(`{${"id"}}`, encodeURIComponent(String(requestParameters.id))),
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        });

        return new runtime.JSONApiResponse(response, (jsonValue) => EntIllnessFromJSON(jsonValue));
    }

    /**
     * get illness by ID
     * Get a illness entity by ID
     */
    async getIllness(requestParameters: GetIllnessRequest): Promise<EntIllness> {
        const response = await this.getIllnessRaw(requestParameters);
        return await response.value();
    }

    /**
     * get patient by ID
     * Get a patient entity by ID
     */
    async getPatientRaw(requestParameters: GetPatientRequest): Promise<runtime.ApiResponse<EntPatient>> {
        if (requestParameters.id === null || requestParameters.id === undefined) {
            throw new runtime.RequiredError('id','Required parameter requestParameters.id was null or undefined when calling getPatient.');
        }

        const queryParameters: runtime.HTTPQuery = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/patients/{id}`.replace(`{${"id"}}`, encodeURIComponent(String(requestParameters.id))),
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        });

        return new runtime.JSONApiResponse(response, (jsonValue) => EntPatientFromJSON(jsonValue));
    }

    /**
     * get patient by ID
     * Get a patient entity by ID
     */
    async getPatient(requestParameters: GetPatientRequest): Promise<EntPatient> {
        const response = await this.getPatientRaw(requestParameters);
        return await response.value();
    }

    /**
     * get symptoms by ID
     * Get a symptoms entity by ID
     */
    async getSymptomsRaw(requestParameters: GetSymptomsRequest): Promise<runtime.ApiResponse<EntSymptoms>> {
        if (requestParameters.id === null || requestParameters.id === undefined) {
            throw new runtime.RequiredError('id','Required parameter requestParameters.id was null or undefined when calling getSymptoms.');
        }

        const queryParameters: runtime.HTTPQuery = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/symptomss/{id}`.replace(`{${"id"}}`, encodeURIComponent(String(requestParameters.id))),
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        });

        return new runtime.JSONApiResponse(response, (jsonValue) => EntSymptomsFromJSON(jsonValue));
    }

    /**
     * get symptoms by ID
     * Get a symptoms entity by ID
     */
    async getSymptoms(requestParameters: GetSymptomsRequest): Promise<EntSymptoms> {
        const response = await this.getSymptomsRaw(requestParameters);
        return await response.value();
    }

    /**
     * list Diagnose entities
     * List Diagnose entities
     */
    async listDiagnoseRaw(requestParameters: ListDiagnoseRequest): Promise<runtime.ApiResponse<Array<EntDiagnose>>> {
        const queryParameters: runtime.HTTPQuery = {};

        if (requestParameters.limit !== undefined) {
            queryParameters['limit'] = requestParameters.limit;
        }

        if (requestParameters.offset !== undefined) {
            queryParameters['offset'] = requestParameters.offset;
        }

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/diagnoses`,
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        });

        return new runtime.JSONApiResponse(response, (jsonValue) => jsonValue.map(EntDiagnoseFromJSON));
    }

    /**
     * list Diagnose entities
     * List Diagnose entities
     */
    async listDiagnose(requestParameters: ListDiagnoseRequest): Promise<Array<EntDiagnose>> {
        const response = await this.listDiagnoseRaw(requestParameters);
        return await response.value();
    }

    /**
     * list doctor entities
     * List doctor entities
     */
    async listDoctorRaw(requestParameters: ListDoctorRequest): Promise<runtime.ApiResponse<Array<EntDoctor>>> {
        const queryParameters: runtime.HTTPQuery = {};

        if (requestParameters.limit !== undefined) {
            queryParameters['limit'] = requestParameters.limit;
        }

        if (requestParameters.offset !== undefined) {
            queryParameters['offset'] = requestParameters.offset;
        }

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/doctors`,
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        });

        return new runtime.JSONApiResponse(response, (jsonValue) => jsonValue.map(EntDoctorFromJSON));
    }

    /**
     * list doctor entities
     * List doctor entities
     */
    async listDoctor(requestParameters: ListDoctorRequest): Promise<Array<EntDoctor>> {
        const response = await this.listDoctorRaw(requestParameters);
        return await response.value();
    }

    /**
     * list illness entities
     * List illness entities
     */
    async listIllnessRaw(requestParameters: ListIllnessRequest): Promise<runtime.ApiResponse<Array<EntIllness>>> {
        const queryParameters: runtime.HTTPQuery = {};

        if (requestParameters.limit !== undefined) {
            queryParameters['limit'] = requestParameters.limit;
        }

        if (requestParameters.offset !== undefined) {
            queryParameters['offset'] = requestParameters.offset;
        }

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/illnesss`,
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        });

        return new runtime.JSONApiResponse(response, (jsonValue) => jsonValue.map(EntIllnessFromJSON));
    }

    /**
     * list illness entities
     * List illness entities
     */
    async listIllness(requestParameters: ListIllnessRequest): Promise<Array<EntIllness>> {
        const response = await this.listIllnessRaw(requestParameters);
        return await response.value();
    }

    /**
     * list patient entities
     * List patient entities
     */
    async listPatientRaw(requestParameters: ListPatientRequest): Promise<runtime.ApiResponse<Array<EntPatient>>> {
        const queryParameters: runtime.HTTPQuery = {};

        if (requestParameters.limit !== undefined) {
            queryParameters['limit'] = requestParameters.limit;
        }

        if (requestParameters.offset !== undefined) {
            queryParameters['offset'] = requestParameters.offset;
        }

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/patients`,
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        });

        return new runtime.JSONApiResponse(response, (jsonValue) => jsonValue.map(EntPatientFromJSON));
    }

    /**
     * list patient entities
     * List patient entities
     */
    async listPatient(requestParameters: ListPatientRequest): Promise<Array<EntPatient>> {
        const response = await this.listPatientRaw(requestParameters);
        return await response.value();
    }

    /**
     * list symptoms entities
     * List symptoms entities
     */
    async listSymptomsRaw(requestParameters: ListSymptomsRequest): Promise<runtime.ApiResponse<Array<EntSymptoms>>> {
        const queryParameters: runtime.HTTPQuery = {};

        if (requestParameters.limit !== undefined) {
            queryParameters['limit'] = requestParameters.limit;
        }

        if (requestParameters.offset !== undefined) {
            queryParameters['offset'] = requestParameters.offset;
        }

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/symptomss`,
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        });

        return new runtime.JSONApiResponse(response, (jsonValue) => jsonValue.map(EntSymptomsFromJSON));
    }

    /**
     * list symptoms entities
     * List symptoms entities
     */
    async listSymptoms(requestParameters: ListSymptomsRequest): Promise<Array<EntSymptoms>> {
        const response = await this.listSymptomsRaw(requestParameters);
        return await response.value();
    }

    /**
     * update diagnose by ID
     * Update a diagnose entity by ID
     */
    async updateDiagnoseRaw(requestParameters: UpdateDiagnoseRequest): Promise<runtime.ApiResponse<EntDiagnose>> {
        if (requestParameters.id === null || requestParameters.id === undefined) {
            throw new runtime.RequiredError('id','Required parameter requestParameters.id was null or undefined when calling updateDiagnose.');
        }

        if (requestParameters.drugAllergy === null || requestParameters.drugAllergy === undefined) {
            throw new runtime.RequiredError('drugAllergy','Required parameter requestParameters.drugAllergy was null or undefined when calling updateDiagnose.');
        }

        const queryParameters: runtime.HTTPQuery = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        const response = await this.request({
            path: `/diagnoses/{id}`.replace(`{${"id"}}`, encodeURIComponent(String(requestParameters.id))),
            method: 'PUT',
            headers: headerParameters,
            query: queryParameters,
            body: EntDiagnoseToJSON(requestParameters.drugAllergy),
        });

        return new runtime.JSONApiResponse(response, (jsonValue) => EntDiagnoseFromJSON(jsonValue));
    }

    /**
     * update diagnose by ID
     * Update a diagnose entity by ID
     */
    async updateDiagnose(requestParameters: UpdateDiagnoseRequest): Promise<EntDiagnose> {
        const response = await this.updateDiagnoseRaw(requestParameters);
        return await response.value();
    }

    /**
     * update doctor by ID
     * Update a doctor entity by ID
     */
    async updateDoctorRaw(requestParameters: UpdateDoctorRequest): Promise<runtime.ApiResponse<EntDoctor>> {
        if (requestParameters.id === null || requestParameters.id === undefined) {
            throw new runtime.RequiredError('id','Required parameter requestParameters.id was null or undefined when calling updateDoctor.');
        }

        if (requestParameters.doctor === null || requestParameters.doctor === undefined) {
            throw new runtime.RequiredError('doctor','Required parameter requestParameters.doctor was null or undefined when calling updateDoctor.');
        }

        const queryParameters: runtime.HTTPQuery = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        const response = await this.request({
            path: `/doctors/{id}`.replace(`{${"id"}}`, encodeURIComponent(String(requestParameters.id))),
            method: 'PUT',
            headers: headerParameters,
            query: queryParameters,
            body: EntDoctorToJSON(requestParameters.doctor),
        });

        return new runtime.JSONApiResponse(response, (jsonValue) => EntDoctorFromJSON(jsonValue));
    }

    /**
     * update doctor by ID
     * Update a doctor entity by ID
     */
    async updateDoctor(requestParameters: UpdateDoctorRequest): Promise<EntDoctor> {
        const response = await this.updateDoctorRaw(requestParameters);
        return await response.value();
    }

    /**
     * update illness by ID
     * Update a illness entity by ID
     */
    async updateIllnessRaw(requestParameters: UpdateIllnessRequest): Promise<runtime.ApiResponse<EntIllness>> {
        if (requestParameters.id === null || requestParameters.id === undefined) {
            throw new runtime.RequiredError('id','Required parameter requestParameters.id was null or undefined when calling updateIllness.');
        }

        if (requestParameters.illness === null || requestParameters.illness === undefined) {
            throw new runtime.RequiredError('illness','Required parameter requestParameters.illness was null or undefined when calling updateIllness.');
        }

        const queryParameters: runtime.HTTPQuery = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        const response = await this.request({
            path: `/illnesss/{id}`.replace(`{${"id"}}`, encodeURIComponent(String(requestParameters.id))),
            method: 'PUT',
            headers: headerParameters,
            query: queryParameters,
            body: EntIllnessToJSON(requestParameters.illness),
        });

        return new runtime.JSONApiResponse(response, (jsonValue) => EntIllnessFromJSON(jsonValue));
    }

    /**
     * update illness by ID
     * Update a illness entity by ID
     */
    async updateIllness(requestParameters: UpdateIllnessRequest): Promise<EntIllness> {
        const response = await this.updateIllnessRaw(requestParameters);
        return await response.value();
    }

    /**
     * update patient by ID
     * Update a patient entity by ID
     */
    async updatePatientRaw(requestParameters: UpdatePatientRequest): Promise<runtime.ApiResponse<EntPatient>> {
        if (requestParameters.id === null || requestParameters.id === undefined) {
            throw new runtime.RequiredError('id','Required parameter requestParameters.id was null or undefined when calling updatePatient.');
        }

        if (requestParameters.patient === null || requestParameters.patient === undefined) {
            throw new runtime.RequiredError('patient','Required parameter requestParameters.patient was null or undefined when calling updatePatient.');
        }

        const queryParameters: runtime.HTTPQuery = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        const response = await this.request({
            path: `/patients/{id}`.replace(`{${"id"}}`, encodeURIComponent(String(requestParameters.id))),
            method: 'PUT',
            headers: headerParameters,
            query: queryParameters,
            body: EntPatientToJSON(requestParameters.patient),
        });

        return new runtime.JSONApiResponse(response, (jsonValue) => EntPatientFromJSON(jsonValue));
    }

    /**
     * update patient by ID
     * Update a patient entity by ID
     */
    async updatePatient(requestParameters: UpdatePatientRequest): Promise<EntPatient> {
        const response = await this.updatePatientRaw(requestParameters);
        return await response.value();
    }

    /**
     * update symptoms by ID
     * Update a symptoms entity by ID
     */
    async updateSymptomsRaw(requestParameters: UpdateSymptomsRequest): Promise<runtime.ApiResponse<EntSymptoms>> {
        if (requestParameters.id === null || requestParameters.id === undefined) {
            throw new runtime.RequiredError('id','Required parameter requestParameters.id was null or undefined when calling updateSymptoms.');
        }

        if (requestParameters.symptoms === null || requestParameters.symptoms === undefined) {
            throw new runtime.RequiredError('symptoms','Required parameter requestParameters.symptoms was null or undefined when calling updateSymptoms.');
        }

        const queryParameters: runtime.HTTPQuery = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        const response = await this.request({
            path: `/symptomss/{id}`.replace(`{${"id"}}`, encodeURIComponent(String(requestParameters.id))),
            method: 'PUT',
            headers: headerParameters,
            query: queryParameters,
            body: EntSymptomsToJSON(requestParameters.symptoms),
        });

        return new runtime.JSONApiResponse(response, (jsonValue) => EntSymptomsFromJSON(jsonValue));
    }

    /**
     * update symptoms by ID
     * Update a symptoms entity by ID
     */
    async updateSymptoms(requestParameters: UpdateSymptomsRequest): Promise<EntSymptoms> {
        const response = await this.updateSymptomsRaw(requestParameters);
        return await response.value();
    }

}
