import {useEffect, useState} from "react";
import axios from "axios";
import {Container, Form, Spinner} from "react-bootstrap";
import {useNavigate} from "react-router-dom";


export const Home = () => {
    const navigate = useNavigate()
    const [error, setErr] = useState('')
    const [loading, isLoading] = useState(true)
    const [data, setData] = useState([])
    const [link, setLink] = useState('')

    const submitHandler = async e => {
        e.preventDefault()
        await axios.post('https://url-shortener-kjie.onrender.com/api/url', {
            origin: link,
        }, {withCredentials: true, headers: {Authorization: `Bearer ${localStorage.getItem("token")}`}})
            .then(function (response) {
                if (response.status == 201) {
                    setData([...data, response.data])
                    setLink('')
                }

            })
    }

    useEffect(() => {
        axios.get('https://url-shortener-kjie.onrender.com/api/url', {
            headers: {Authorization: `Bearer ${localStorage.getItem("token")}`}
        }, {withCredentials: true})
            .then(function (response) {
                if (response.data == null) {
                    setData([])
                    isLoading(false)
                } else {
                    setData(response.data)
                    isLoading(false)
                }
            })
            .catch(function (error) {
                navigate('/signin')
                setErr(error)
            })
    }, []);

    return (
        <Container>
            <Form method="POST" onSubmit={submitHandler}>
                <Form.Label htmlFor="inputPassword5">URL Address</Form.Label>
                <Form.Control
                    type="text"
                    id="inputPassword5"
                    aria-describedby="passwordHelpBlock"
                    placeholder="Например vk.com"
                    value={link}
                    onChange={e => setLink(e.target.value)}
                />
                <Form.Text id="passwordHelpBlock" muted>
                    Напишите адрес, который вы хотите сократить
                </Form.Text>
            </Form>
            <h4 style={{fontWeight: `bold`, marginTop: `30px`}}>Ваши короткие ссылки</h4>
            {loading ? (
                <Spinner animation="border" role="status" style={{position: `absolute`, left: `50%`}}>
                    <span className="visually-hidden">Loading...</span>
                </Spinner>
            ) : (
                data ? (
                    <ul className="list-group">
                        {data.map(url => <li className="list-group-item mb-20"><a target="_blank" style={{textDecoration: `none`}} className="text-muted" href={"https://url-shortener-kjie.onrender.com/"+url.UUID}>https://url-shortener-kjie.onrender.com/{url.UUID}</a>, {url['origin']}</li>)}
                    </ul>
                ) : (
                    <h1></h1>
                )
            )}
        </Container>
    )

}