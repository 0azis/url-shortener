import {redirect, useParams} from "react-router-dom";
import axios from "axios";


export const Url = () => {
    let { uuid } = useParams()
    const shortLink = async() => {
        await axios.get(`http://localhost:8080/api/${uuid}`, {
            headers: {Authorization: `Bearer ${localStorage.getItem("token")}`}
        }, {withCredentials: true})
            .then(function (response) {
                console.log(response.data.message)
                window.location.replace(`https://${response.data.message}`);
            })
    }
    shortLink()
}
