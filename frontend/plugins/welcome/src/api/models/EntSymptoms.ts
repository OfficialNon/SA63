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
    EntSymptomsEdges,
    EntSymptomsEdgesFromJSON,
    EntSymptomsEdgesFromJSONTyped,
    EntSymptomsEdgesToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntSymptoms
 */
export interface EntSymptoms {
    /**
     * Manner holds the value of the "Manner" field.
     * @type {string}
     * @memberof EntSymptoms
     */
    manner?: string;
    /**
     * 
     * @type {EntSymptomsEdges}
     * @memberof EntSymptoms
     */
    edges?: EntSymptomsEdges;
    /**
     * ID of the ent.
     * @type {number}
     * @memberof EntSymptoms
     */
    id?: number;
}

export function EntSymptomsFromJSON(json: any): EntSymptoms {
    return EntSymptomsFromJSONTyped(json, false);
}

export function EntSymptomsFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntSymptoms {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'manner': !exists(json, 'Manner') ? undefined : json['Manner'],
        'edges': !exists(json, 'edges') ? undefined : EntSymptomsEdgesFromJSON(json['edges']),
        'id': !exists(json, 'id') ? undefined : json['id'],
    };
}

export function EntSymptomsToJSON(value?: EntSymptoms | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'Manner': value.manner,
        'edges': EntSymptomsEdgesToJSON(value.edges),
        'id': value.id,
    };
}


