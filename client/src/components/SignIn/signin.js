
import { Form, Button, Alert } from "react-bootstrap";
import {useState} from "react";
import {useNavigate} from "react-router-dom";
import { ToastContainer, toast } from 'react-toastify';
import 'react-toastify/dist/ReactToastify.css';
import axios from "axios";
export const SignIn = () => {
    const notify = (text) => toast.error(text, {
        position: "top-right",
        autoClose: 5000,
        hideProgressBar: false,
        closeOnClick: true,
        pauseOnHover: true,
        draggable: true,
        progress: undefined,
        theme: "colored",
    });
    const navigate = useNavigate();
    const [loading, isLoading] = useState(false)
    const [credentials, setCredentials] = useState({email: '', password: ''})

    const submitHandler = async e => {
        e.preventDefault()
        isLoading(true)
        await axios.post('http://localhost:8080/api/auth/signin', {
            email: credentials.email,
            password: credentials.password
        }, {withCredentials: true})
            .then(function(response) {
                if (response.status === 200) {
                    localStorage.setItem("token", response.data['message'])
                    navigate('/home')
                }

            })
            .catch(function (err) {
                isLoading(false)
                notify(err.response.data.message)
            })
    }


    return (
        <>
        <ToastContainer
            position="top-right"
            autoClose={5000}
            hideProgressBar={false}
            newestOnTop={false}
            closeOnClick
            rtl={false}
            pauseOnFocusLoss
            draggable
            pauseOnHover
            theme="colored"
        />

            <div
                className="sign-in__wrapper"

                style={{ width: `450px`, height: `auto`, position: `absolute`, top: `50%`, left: `50%`, transform: `translate(-50%, -50%)`}}
            >
                {/* Overlay */}
                <div className="sign-in__backdrop"></div>
                {/* Form */}
                <Form className="shadow p-4 bg-white rounded" method="POST" onSubmit={submitHandler}>
                    {/* Header */}
                    <div className="h4 mb-2 text-center">Вход</div>
                    <Form.Group className="mb-2" controlId="email">
                        <Form.Label>Email</Form.Label>
                        <Form.Control
                            type="email"
                            placeholder="Email"
                            value={credentials.email}
                            onChange={e => setCredentials({ ...credentials, email: e.target.value })}
                            required
                        />
                    </Form.Group>
                    <Form.Group className="mb-2" controlId="password">
                        <Form.Label>Password</Form.Label>
                        <Form.Control
                            type="password"
                            placeholder="Password"
                            value={credentials.password}
                            onChange={e => setCredentials({ ...credentials, password: e.target.value })}
                            required
                        />
                    </Form.Group>
                    {!loading ? (
                        <Button className="w-100" variant="primary" type="submit" onClick={notify}>
                        Вход
                        </Button>
                    ) : (
                        <Button className="w-100" variant="primary" type="submit">
                            Входим..
                        </Button>
                    )}


                    <div className="d-grid justify-content-end">
                        <a
                            href="/signup"
                            className="text-muted"
                            variant="link"
                            style = {{marginTop: `10px`}}
                        >
                            Регистрация
                        </a>
                    </div>
                </Form>

            </div>
        </>
    );
}