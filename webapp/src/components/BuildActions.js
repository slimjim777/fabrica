import React from 'react';
import {T} from "./Utils";
import {Link} from "@canonical/react-components"

function BuildActions(props) {
    return (
        <div>
            <Link href={'/builds/'+props.id}  title={T("view")}>
                <img className="action"  src="/static/images/show.svg" alt={T("view")}/>
            </Link>
            {props.download ?
                <Link href={'/v1/builds/' + props.id + '/download'} title={T("download")}>
                    <img className="action" src="/static/images/download.svg" alt={T("download")}/>
                </Link>
                :
                ''
            }
            <Link href="" title={T("delete")} onClick={props.onConfirmDelete}>
                <img className="action" src="/static/images/delete.svg" alt={T("delete")} data-key={props.id}/>
            </Link>
        </div>
    );
}

export default BuildActions;