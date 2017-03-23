import  mysql.connector as mariadb
import json
import datetime
from flask import Flask
from flask_restful import reqparse, abort, Api, Resource
from flask_jwt import JWT, jwt_required, current_identity
from werkzeug.security import safe_str_cmp

app = Flask(__name__)
api = Api(app)

def authenticate(username, password):


def identity(payload):


class getCitizen_ID(Resource):
    def get(self, ID):
    	resp = getCitizen_D_ID(ID).replace("\"","'")
		return resp

class getCitizen_Name(Resource):
	def get(self, Name):
		return getCitizen_D_Name(Name).replace("\"","'")

class getCitizen_Dist_City(Resource):
	def get(self, City_Dist):
		return getCitizen_D_Dist_City(City_Dist).replace("\"","'")

class getCitizen_Adhaar(Resource):
	def get(self, Adhaar):
		return getCitizen_D_Adhaar(Adhaar).replace("\"","'")

class getCitizen_Position(Resource):
	def get(self, Position):
		return getCitizen_D_Position(Position).replace("\"","'")

class getCitizen_Position(Resource):
	def get(self, Position):
		return getCitizen_D_Position(Position).replace("\"","'")

api.add_resource(getCitizen_ID, '/Citizen/<ID>')


app.config['SECRET_KEY'] = 'super-secret'

jwt = JWT(app, authenticate, identity)

if __name__ == '__main__':
	app.run(host= '0.0.0.0',port=5000,debug = True)
