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
    EntProblem,
    EntProblemFromJSON,
    EntProblemFromJSONTyped,
    EntProblemToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntProblemTitleEdges
 */
export interface EntProblemTitleEdges {
    /**
     * Problemtitle2problem holds the value of the problemtitle2problem edge.
     * @type {Array<EntProblem>}
     * @memberof EntProblemTitleEdges
     */
    problemtitle2problem?: Array<EntProblem>;
}

export function EntProblemTitleEdgesFromJSON(json: any): EntProblemTitleEdges {
    return EntProblemTitleEdgesFromJSONTyped(json, false);
}

export function EntProblemTitleEdgesFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntProblemTitleEdges {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'problemtitle2problem': !exists(json, 'problemtitle2problem') ? undefined : ((json['problemtitle2problem'] as Array<any>).map(EntProblemFromJSON)),
    };
}

export function EntProblemTitleEdgesToJSON(value?: EntProblemTitleEdges | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'problemtitle2problem': value.problemtitle2problem === undefined ? undefined : ((value.problemtitle2problem as Array<any>).map(EntProblemToJSON)),
    };
}


