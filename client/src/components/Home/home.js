import {useEffect, useState} from "react";
import axios from "axios";
import {Button, Container, Form, Spinner, Stack} from "react-bootstrap";
import {useNavigate} from "react-router-dom";


export const Home = () => {
    const navigate = useNavigate()
    const [error, setErr] = useState('')
    const [loading, isLoading] = useState(true)
    const [data, setData] = useState([])
    const [link, setLink] = useState('')
    const [delUrl, setDelUrl] = useState('')

    const deleteHandler = async e => {
        e.preventDefault()
        await axios.delete(`http://localhost:8080/api/url?url_id=${delUrl}`, {
            headers: {Authorization: `Bearer ${localStorage.getItem("token")}`}
        }, {withCredentials: true})
            .then(function (response) {
                if (response.status == 200) {
                    setData(current =>
                        current.filter(url => {
                            return url.UUID != delUrl
                        }),
                    );
                }

            })
    }

    const submitHandler = async e => {
        e.preventDefault()
        await axios.post('http://localhost:8080/api/url', {
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
        axios.get('http://localhost:8080/api/url', {
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
                        {data.map(url =>
                            <li className="list-group-item mb-20" style={{display: `flex`, alignItems: `center`, justifyContent: `space-between`}}>
                                <div className="info">
                                    <a target="_blank" style={{textDecoration: `none`}} className="text-muted" href={"http://localhost:3000/"+url.UUID}>url-shortener-dusky-zeta.vercel.app/{url.UUID}</a>, {url['origin']}
                                </div>
                                <form method="DELETE" onSubmit={deleteHandler}>
                                    <Button variant="outline-danger" style={{fontSize: `14px`}} type="submit" onClick={e => setDelUrl(url.UUID)}>Удалить</Button>
                                </form>
                            </li>
                        )}
                    </ul>
                ) : (
                    <h1></h1>
                )
            )}
        </Container>
    )

}