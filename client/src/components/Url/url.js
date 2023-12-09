import {redirect, useParams} from "react-router-dom";
import axios from "axios";


export const Url = () => {
    let { uuid } = useParams()
    const shortLink = async() => {
        await axios.get(`https://url-shortener-kjie.onrender.com/api/url/${uuid}`, {
            headers: {Authorization: `Bearer ${localStorage.getItem("token")}`}
        }, {withCredentials: true})
            .then(function (response) {
                console.log(response.data.message)
                window.location.replace(`https://${response.data.message}`);
            })
    }
    shortLink()
}
