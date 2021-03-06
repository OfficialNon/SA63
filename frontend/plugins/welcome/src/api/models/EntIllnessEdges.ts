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

import { exists, mapValues } from '../runtime';
import {
    EntDiagnose,
    EntDiagnoseFromJSON,
    EntDiagnoseFromJSONTyped,
    EntDiagnoseToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntIllnessEdges
 */
export interface EntIllnessEdges {
    /**
     * IllnessDiagnose holds the value of the illness_diagnose edge.
     * @type {Array<EntDiagnose>}
     * @memberof EntIllnessEdges
     */
    illnessDiagnose?: Array<EntDiagnose>;
}

export function EntIllnessEdgesFromJSON(json: any): EntIllnessEdges {
    return EntIllnessEdgesFromJSONTyped(json, false);
}

export function EntIllnessEdgesFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntIllnessEdges {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'illnessDiagnose': !exists(json, 'illnessDiagnose') ? undefined : ((json['illnessDiagnose'] as Array<any>).map(EntDiagnoseFromJSON)),
    };
}

export function EntIllnessEdgesToJSON(value?: EntIllnessEdges | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'illnessDiagnose': value.illnessDiagnose === undefined ? undefined : ((value.illnessDiagnose as Array<any>).map(EntDiagnoseToJSON)),
    };
}


